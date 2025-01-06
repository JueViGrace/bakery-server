package api

import (
	"errors"
	"strings"
	"time"

	"github.com/JueViGrace/bakery-go/internal/data"
	"github.com/JueViGrace/bakery-go/internal/types"
	"github.com/JueViGrace/bakery-go/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthData struct {
	jwt  JwtData
	role string
}

type JwtData struct {
	token  *jwt.Token
	claims util.JWTClaims
}

func (a *api) adminAuthMiddleware(c *fiber.Ctx) error {
	data, err := getUserDataForReq(c, a.db)
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	if data.role != types.Admin {
		res := types.RespondForbbiden("permission denied", "Failed")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) sessionMiddleware(c *fiber.Ctx) error {
	data, err := getUserDataForReq(c, a.db)
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	_, err = a.db.SessionStore().GetTokenById(data.jwt.claims.UserId.String())
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) userIdMiddleware(c *fiber.Ctx) error {
	data, err := getUserDataForReq(c, a.db)
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	if data.role == types.Admin {
		return c.Next()
	}

	id, err := util.GetIdFromParams(c.Params("id"))
	if err != nil {
		res := types.RespondBadRequest(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	if data.jwt.claims.UserId != *id {
		res := types.RespondForbbiden("Forbbiden action", "Failed")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

// todo: orders middleware

func getUserDataForReq(c *fiber.Ctx, db data.Storage) (*AuthData, error) {
	jwt, err := extractJWTFromHeader(c, func(s string) {
		db.SessionStore().DeleteTokenByToken(s)
	})
	if err != nil {
		return nil, err
	}

	user, err := db.UserStore().GetUserById(&jwt.claims.UserId)
	if err != nil {
		return nil, err
	}

	return &AuthData{
		jwt:  *jwt,
		role: user.Role,
	}, nil
}

func extractJWTFromHeader(c *fiber.Ctx, expired func(string)) (*JwtData, error) {
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
		expired(tokenString)
		return nil, errors.New("permission denied")
	}

	claims, ok := token.Claims.(util.JWTClaims)
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

	if claims.Issuer != util.Issuer {
		return nil, errors.New("permision denied")
	}

	return &JwtData{
		token:  token,
		claims: claims,
	}, nil
}
