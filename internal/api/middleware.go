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

func (a *api) adminAuthMiddleware(c *fiber.Ctx) error {
	token, err := a.extractJWTFromHeader(c)
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	if claims, ok := token.Claims.(util.JWTClaims); !ok || claims.Role != "admin" {
		res := types.RespondForbbiden("permission denied", "Failed")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) authMiddleware(c *fiber.Ctx) error {
	_, err := a.extractJWTFromHeader(c)
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) checkUserIdParamMiddleware(c *fiber.Ctx) error {
	id, err := util.GetIdFromParams(c.Params("id"))
	if err != nil {
		res := types.RespondBadRequest(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	token, err := a.extractJWTFromHeader(c)
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	claims := token.Claims.(util.JWTClaims)
	if claims.Role == "admin" {
		return c.Next()
	}

	tokenID, err := util.GetIdFromParams(claims.ID)
	if err != nil {
		res := types.RespondUnauthorized("permission denied", "Failed")
		return c.Status(res.Status).JSON(res)
	}

	if tokenID != id {
		res := types.RespondForbbiden("forbidden resource", "Failed")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) checkUpdateUserMiddleware(c *fiber.Ctx) error {
	ur := new(types.UpdateUserRequest)

	if err := c.BodyParser(ur); err != nil {
		res := types.RespondBadRequest("permission denied", "Failed")
		return c.Status(res.Status).JSON(res)
	}

	token, err := a.extractJWTFromHeader(c)
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	claims := token.Claims.(util.JWTClaims)
	if claims.Role == "admin" {
		return c.Next()
	}

	tokenID, err := util.GetIdFromParams(claims.ID)
	if err != nil {
		res := types.RespondUnauthorized("permission denied", "Failed")
		return c.Status(res.Status).JSON(res)
	}

	if tokenID != &ur.ID {
		res := types.RespondForbbiden("forbidden resource", "Failed")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) extractJWTFromHeader(c *fiber.Ctx) (*jwt.Token, error) {
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

	return token, nil
}
