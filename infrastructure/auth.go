package infrastructure

import (
	"android-service/model"
	"errors"
	"io"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type claims struct {
	Email       string
	DisplayName string
	jwt.StandardClaims
}

func GetPublicKey() []byte {
	file, err := os.Open("crypt/public.pem")
	if err != nil {
		return nil
	}
	defer file.Close()
	fileByte, _ := io.ReadAll(file)
	return fileByte
}

func CreateJwt(model model.User) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
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

func GetMiddleWareConfig() echojwt.Config {
	return echojwt.Config{
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				keyBys := GetPublicKey()
				publicKey, _ := jwt.ParseRSAPublicKeyFromPEM(keyBys)
				return publicKey, nil
			}
			token, err := jwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/user/login" {
				return true
			}
			if c.Path() == "/user" {
				return true
			}
			if strings.HasPrefix(c.Path(), "/reset") {
				return true
			}
			return false
		},
	}
}
