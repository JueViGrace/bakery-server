package util

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
)

type JwtData struct {
	Token  *jwt.Token
	Claims *JWTClaims
}

type userClaims struct {
	UserId   uuid.UUID `json:"userId"`
	FullName string    `json:"fullName"`
}

type JWTClaims struct {
	userClaims
	jwt.RegisteredClaims
}

const (
	Issuer string = "BakeryServer"
)

var (
	jwtSecret string           = os.Getenv("JWT_SECRET")
	Audience  jwt.ClaimStrings = jwt.ClaimStrings{
		"api",
	}
)

// TODO: make refresh token

func CreateAccessToken(id, fullName string) (string, error) {
	return CreateJWT(id, fullName, time.Now().UTC().Add(1*time.Hour))
}

func CreateRefreshToken(id, fullName string) (string, error) {
	return CreateJWT(id, fullName, time.Now().UTC().Add(24*time.Hour))
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
	return jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
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
		log.Error(err.Error())
		expired(tokenString)
		return nil, errors.New("permission denied")
	}

	claims, ok := token.Claims.(JWTClaims)
	if !ok || !token.Valid {
		log.Error("invalid claims or expired")
		expired(tokenString)
		return nil, errors.New("permission denied")
	}

	if len(claims.Audience) > 1 || claims.
		Audience[0] != "api" {
		log.Error("bad audience: %v", claims.Audience)
		expired(tokenString)
		return nil, errors.New("permision denied")
	}

	if claims.Issuer != Issuer {
		log.Error("bad issuer: %v", claims.Issuer)
		expired(tokenString)
		return nil, errors.New("permision denied")
	}

	return &JwtData{
		Token:  token,
		Claims: &claims,
	}, nil
}
