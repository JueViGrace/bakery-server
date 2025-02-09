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
	UserId    uuid.UUID `json:"user_id"`
	SessionID uuid.UUID `json:"session_id"`
	Username  string    `json:"username"`
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

func CreateAccessToken(userId, sessionId uuid.UUID, username string) (string, error) {
	var accessExpiration time.Time = time.Now().UTC().Add(1 * time.Hour)
	return CreateJWT(userId, sessionId, username, accessExpiration)
}

func CreateRefreshToken(userId, sessionId uuid.UUID, username string) (string, error) {
	var refreshExpiration time.Time = time.Now().UTC().Add(24 * time.Hour)
	return CreateJWT(userId, sessionId, username, refreshExpiration)
}

func CreateJWT(userId, sessionId uuid.UUID, username string, expiration time.Time) (string, error) {
	tokenId, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	claims := JWTClaims{
		userClaims{
			UserId:    userId,
			SessionID: sessionId,
			Username:  username,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
			Issuer:    Issuer,
			Subject:   userId.String(),
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

	claims, ok := token.Claims.(*JWTClaims)
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
		Claims: claims,
	}, nil
}
