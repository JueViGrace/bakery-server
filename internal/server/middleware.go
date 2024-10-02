package server

import (
	"errors"
	"strings"
	"time"

	"github.com/JueViGrace/bakery-go/internal/data"
	"github.com/JueViGrace/bakery-go/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (s *FiberServer) adminAuthMiddleware(c *fiber.Ctx) error {
	token, err := s.extractJWTFromHeader(c)
	if err != nil {
		return RespondUnauthorized(c, err.Error(), "Failed")
	}

	if claims, ok := token.Claims.(util.JWTClaims); !ok || claims.Role != "admin" {
		return RespondForbbiden(c, "permission denied", "Failed")
	}

	return c.Next()
}

func (s *FiberServer) authMiddleware(c *fiber.Ctx) error {
	_, err := s.extractJWTFromHeader(c)
	if err != nil {
		return RespondUnauthorized(c, err.Error(), "Failed")
	}

	return c.Next()
}

func (s *FiberServer) checkUserIdParamMiddleware(c *fiber.Ctx) error {
	id, err := util.GetIdFromParams(c.Params("id"))
	if err != nil {
		return RespondBadRequest(c, err.Error(), "Failed")
	}

	token, err := s.extractJWTFromHeader(c)
	if err != nil {
		return RespondUnauthorized(c, err.Error(), "Failed")
	}

	claims := token.Claims.(util.JWTClaims)
	if claims.Role == "admin" {
		return c.Next()
	}

	tokenID, err := util.GetIdFromParams(claims.ID)
	if err != nil {
		return RespondUnauthorized(c, "permission denied", "Failed")
	}

	if tokenID != id {
		return RespondForbbiden(c, "forbidden resource", "Failed")
	}

	return c.Next()
}

func (s *FiberServer) checkUpdateUserMiddleware(c *fiber.Ctx) error {
	ur := new(data.UpdateUserRequest)

	if err := c.BodyParser(ur); err != nil {
		return RespondBadRequest(c, "permission denied", "Failed")
	}

	token, err := s.extractJWTFromHeader(c)
	if err != nil {
		return RespondUnauthorized(c, err.Error(), "Failed")
	}

	claims := token.Claims.(util.JWTClaims)
	if claims.Role == "admin" {
		return c.Next()
	}

	tokenID, err := util.GetIdFromParams(claims.ID)
	if err != nil {
		return RespondUnauthorized(c, "permission denied", "Failed")
	}

	if tokenID != &ur.ID {
		return RespondForbbiden(c, "forbidden resource", "Failed")
	}

	return c.Next()
}

func (s *FiberServer) extractJWTFromHeader(c *fiber.Ctx) (*jwt.Token, error) {
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
