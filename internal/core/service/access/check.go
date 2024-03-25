package access

import (
	"context"
	"fmt"
	"strings"

	modelAccess "github.com/pillarion/practice-auth/internal/core/model/access"
	modelJounal "github.com/pillarion/practice-auth/internal/core/model/journal"
	jwtTool "github.com/pillarion/practice-auth/internal/core/tools/jwt"
	"google.golang.org/grpc/metadata"
)

const authPrefix = "Bearer "

func (s *service) Check(ctx context.Context, endpoint string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return fmt.Errorf("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return fmt.Errorf("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	claims, err := jwtTool.VerifyJWT(accessToken, []byte(s.jwtConfig.Secret), jwtTool.JWTAccessTokenVariant)
	if err != nil {
		return err
	}

	var matrix []modelAccess.Access
	err = s.txManager.ReadCommitted(
		ctx,
		func(ctx context.Context) error {
			var errTx error
			matrix, errTx = s.accessRepo.AccessMatrix(ctx, endpoint)
			if errTx != nil {
				fmt.Println(errTx)
				return errTx
			}

			_, errTx = s.journalRepo.Insert(ctx, &modelJounal.Journal{
				Action: "Check endpoint",
			})
			if errTx != nil {
				return errTx
			}

			return nil
		})
	if err != nil {
		return err
	}

	for _, m := range matrix {
		if m.Role == claims.Role {
			return nil
		}
	}

	return fmt.Errorf("access denied")
}
