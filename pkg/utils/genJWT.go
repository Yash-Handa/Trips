package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenJWTToken generates a JWT token with ID and returns the token and expiry time
func GenJWTToken(ID string) (string, time.Time, error) {
	expiredAt := time.Now().Add(time.Hour * 24 * 7) // a week

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Id:        ID,
		IssuedAt:  time.Now().Unix(),
		Issuer:    "Trips",
	})

	accessToken, err := token.SignedString([]byte(MustGet("JWT_SECRET")))
	if err != nil {
		return "", time.Now(), err
	}

	return accessToken, expiredAt, nil
}
