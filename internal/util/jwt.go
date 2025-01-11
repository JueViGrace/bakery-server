package util

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
)

var (
	jwtSecret string           = os.Getenv("JWT_SECRET")
	Issuer    string           = "BakeryServer"
	Audience  jwt.ClaimStrings = jwt.ClaimStrings{
		"api",
	}
	accessExpiration  time.Time = time.Now().Add(1 * time.Hour)
	refreshExpiration time.Time = time.Now().Add(24 * time.Hour)
)

type JwtData struct {
	Token  *jwt.Token
	Claims JWTClaims
}

type userClaims struct {
	UserId   uuid.UUID `json:"userId"`
	FullName string    `json:"fullName"`
}

type JWTClaims struct {
	userClaims
	jwt.RegisteredClaims
}

// TODO: make refresh token

func CreateAccessToken(id, fullName string) (string, error) {
	return CreateJWT(id, fullName, accessExpiration)
}

func CreateRefreshToken(id, fullName string) (string, error) {
	return CreateJWT(id, fullName, refreshExpiration)
}

func CreateJWT(id, fullName string, expiration time.Time) (string, error) {
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
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    Issuer,
			Subject:   id,
			ID:        tokenId.String(),
			Audience:  Audience,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
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

func ExtractJWTFromHeader(c *fiber.Ctx, expired func(string)) (*JwtData, error) {
	header := strings.Join(c.GetReqHeaders()["Authorization"], "")

	if !strings.HasPrefix(header, "Bearer") {
		return nil, errors.New("permission denied")
	}

	tokenString := strings.Split(header, " ")[1]
	token, err := ValidateJWT(tokenString)
	if err != nil {
		return nil, errors.New("permission denied")
	}

	if !token.Valid {
		expired(tokenString)
		return nil, errors.New("permission denied")
	}

	claims, ok := token.Claims.(JWTClaims)
	if !ok {
		expired(tokenString)
		return nil, errors.New("permission denied")
	}

	if claims.ExpiresAt.Time.UTC().Unix() < time.Now().UTC().Unix() {
		expired(tokenString)
		return nil, errors.New("permision denied")
	}

	if len(claims.Audience) > 1 || claims.
		Audience[0] != "api" {
		return nil, errors.New("permision denied")
	}

	if claims.Issuer != Issuer {
		return nil, errors.New("permision denied")
	}

	return &JwtData{
		Token:  token,
		Claims: claims,
	}, nil
}
