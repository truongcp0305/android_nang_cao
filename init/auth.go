package init

import (
	"android-service/model"
	"io"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type claims struct {
	Email       string
	DisplayName string
	jwt.StandardClaims
}

func GetPublicKey() []byte {
	file, err := os.Open("crypt/private.pem")
	if err != nil {
		return nil
	}
	defer file.Close()
	fileByte, _ := io.ReadAll(file)
	return fileByte
}

func CreateJwt(model model.User) (string, error) {
	expirationTime := time.Now().Add(10 * time.Hour)
	claims := &claims{
		Email:       model.UserName,
		DisplayName: model.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetPublicKey())
}
