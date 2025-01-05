package util

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
)

var jwtSecret string = os.Getenv("JWT_SECRET")

type userClaims struct {
	UserId   uuid.UUID `json:"userId"`
	FullName string    `json:"fullName"`
}

type JWTClaims struct {
	userClaims
	jwt.RegisteredClaims
}

// TODO: make refresh token

func CreateJWT(id, fullName string) (string, error) {
	tokenId, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	userId, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	claims := JWTClaims{
		userClaims{
			UserId:   userId,
			FullName: fullName,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "BakeryServer",
			Subject:   id,
			ID:        tokenId.String(),
			Audience: jwt.ClaimStrings{
				"api",
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	claims := new(JWTClaims)

	return jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
}
