package api

import (
	"errors"
	"strings"
	"time"

	"github.com/JueViGrace/bakery-go/internal/types"
	"github.com/JueViGrace/bakery-go/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthData struct {
	token  *jwt.Token
	claims util.JWTClaims
}

func (a *api) adminAuthMiddleware(c *fiber.Ctx) error {
	data, err := a.extractJWTFromHeader(c)
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	user, err := a.db.UserStore().GetUserById(&data.claims.UserId)

	if user.Role != "admin" {
		res := types.RespondForbbiden("permission denied", "Failed")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) sessionMiddleware(c *fiber.Ctx) error {
	data, err := a.extractJWTFromHeader(c)
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	_, err = a.db.SessionStore().GetTokenById(data.claims.UserId.String())
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) extractJWTFromHeader(c *fiber.Ctx) (*AuthData, error) {
	header := strings.Join(c.GetReqHeaders()["Authorization"], "")

	if !strings.HasPrefix(header, "Bearer") {
		return nil, errors.New("permission denied")
	}

	tokenString := strings.Split(header, " ")[1]
	token, err := util.ValidateJWT(tokenString)
	if err != nil {
		return nil, errors.New("permission denied")
	}

	if !token.Valid {
		return nil, errors.New("permission denied")
	}

	claims, ok := token.Claims.(util.JWTClaims)
	if !ok {
		return nil, errors.New("permission denied")
	}

	if claims.ExpiresAt.Time.UTC().Unix() < time.Now().UTC().Unix() {
		return nil, errors.New("permision denied")
	}

	if len(claims.Audience) > 1 || claims.Audience[0] != "api" {
		return nil, errors.New("permision denied")
	}

	if claims.Issuer != util.Issuer {
		return nil, errors.New("permision denied")
	}

	return &AuthData{
		token:  token,
		claims: claims,
	}, nil
}
