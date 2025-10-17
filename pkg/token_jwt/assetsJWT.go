package tokenjwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID string, secretKey string) (string, error) {
	now := time.Now()
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "mobile-app",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Println(token)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyJWT(tokenStr string, secretKey string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(_ *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
