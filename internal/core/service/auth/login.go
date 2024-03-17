package auth

import (
	"context"
	"fmt"

	modelAuth "github.com/pillarion/practice-auth/internal/core/model/auth"
	modelJounal "github.com/pillarion/practice-auth/internal/core/model/journal"
	modelUser "github.com/pillarion/practice-auth/internal/core/model/user"
	jwtTool "github.com/pillarion/practice-auth/internal/core/tools/jwt"
	passTool "github.com/pillarion/practice-auth/internal/core/tools/password"
)

func (s *service) Login(ctx context.Context, cred modelAuth.Credential) (string, error) {
	var user *modelUser.User
	err := s.txManager.ReadCommitted(
		ctx,
		func(ctx context.Context) error {
			var errTx error

			user, errTx = s.userRepo.SelectByName(ctx, cred.Username)
			if errTx != nil {
				return errTx
			}

			_, errTx = s.journalRepo.Insert(ctx, &modelJounal.Journal{
				Action: "User login",
			})
			if errTx != nil {
				return errTx
			}

			return nil
		})
	if err != nil {
		return "", err
	}

	var refreshToken string
	if passTool.Check(cred.Password, user.Info.Password) {
		refreshToken, err = jwtTool.GenerateJWT(
			user.Info.Name,
			user.Info.Role,
			[]byte(s.jwtConfig.Secret),
			s.jwtConfig.RefreshDuration,
			jwtTool.JWTRefreshTokenVariant,
		)
		if err != nil {
			return "", err
		}

		return refreshToken, nil
	} else {
		return "", fmt.Errorf("invalid credentials")
	}
}
