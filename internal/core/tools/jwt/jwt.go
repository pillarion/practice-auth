package jwt

import (
	"time"

	userModel "github.com/pillarion/practice-auth/internal/core/model/user"
)

func GenerateJWT(user userModel.Info, secret []byte, duration time.Duration) (string, error) {

	return "", nil
}
