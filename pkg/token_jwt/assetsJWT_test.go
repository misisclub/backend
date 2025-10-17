package tokenjwt

import (
	"log"
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

func TestGenerateJWT(t *testing.T) {
	secretKey := "mysecretkey"

	userID := "12345"
	token, err := GenerateJWT(userID, secretKey)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	log.Println(token)
	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(_ *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !parsedToken.Valid {
		t.Fatalf("Expected token to be valid")
	}

	if claims.UserID != userID {
		t.Errorf("Expected userID %v, got %v", userID, claims.UserID)
	}
}
