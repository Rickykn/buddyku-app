package helpers

import (
	"os"
	"strconv"
	"time"

	"github.com/Rickykn/buddyku-app.git/dtos"
	"github.com/golang-jwt/jwt/v4"
)

type IdTokenClaims struct {
	jwt.RegisteredClaims
	User *dtos.ResponseTokenDTO `json:"user"`
}

func CreateJwt(dataToken *dtos.ResponseTokenDTO) (string, error) {

	expired := os.Getenv("EXPIRED")
	expiredInt, _ := strconv.Atoi(expired)
	unixTime := time.Now().Unix()
	tokenExp := unixTime + int64(expiredInt)

	claims := &IdTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: os.Getenv("ISSUER"),
			ExpiresAt: &jwt.NumericDate{
				Time: time.Unix(tokenExp, 0),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
		User: dataToken,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
