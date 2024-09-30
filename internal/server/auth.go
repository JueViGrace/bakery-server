package server

import (
	"github.com/JueViGrace/bakery-go/internal/data"
	"github.com/gofiber/fiber/v2"
)

type AuthRoutes interface {
	SignIn(c *fiber.Ctx) error
	SignUp(c *fiber.Ctx) error
	RecoverPassword(c *fiber.Ctx) error
	ChangeEmail(c *fiber.Ctx) error
}

type authHandler struct {
	au data.AuthStore
}

func NewAuthHandler(au data.AuthStore) AuthRoutes {
	return &authHandler{
		au: au,
	}
}

func (s *FiberServer) AuthRoutes() {
	authGroup := s.Group("/api/auth")

	authHandler := NewAuthHandler(s.db.AuthStore())

	authGroup.Post("/signIn", authHandler.SignIn)
	authGroup.Post("/signUp", authHandler.SignUp)
	authGroup.Post("/recover/password", authHandler.RecoverPassword)
	authGroup.Post("/recover/email", authHandler.ChangeEmail)
}

func (h *authHandler) SignIn(c *fiber.Ctx) error {
	r := new(data.SignInRequest)
	if err := c.BodyParser(r); err != nil {
		return c.JSON(RespondBadRequest(err.Error(), "Failed"))
	}

	token, err := h.au.SignIn(*r)
	if err != nil {
		return c.JSON(RespondNotFound(err.Error(), "Failed"))
	}

	return c.JSON(RespondOk(token, "Success"))
}

func (h *authHandler) SignUp(c *fiber.Ctx) error {
	r := new(data.SignUpRequest)
	if err := c.BodyParser(r); err != nil {
		return c.JSON(RespondBadRequest(err.Error(), "Failed"))
	}

	token, err := h.au.SignUp(*r)
	if err != nil {
		return c.JSON(RespondNotFound(err.Error(), "Failed"))
	}

	return c.JSON(RespondCreated(token, "Success"))
}

func (h *authHandler) RecoverPassword(c *fiber.Ctx) error {
	r := new(data.RecoverPasswordRequest)
	if err := c.BodyParser(r); err != nil {
		return c.JSON(RespondBadRequest(err.Error(), "Failed"))
	}

	msg, err := h.au.RecoverPassword(*r)
	if err != nil {
		return c.JSON(RespondNotFound(err.Error(), "Failed"))
	}

	return c.JSON(RespondAccepted(msg, "Success"))
}

func (h *authHandler) ChangeEmail(c *fiber.Ctx) error {
	r := new(data.ChangeEmailRequest)
	if err := c.BodyParser(r); err != nil {
		return c.JSON(RespondBadRequest(err.Error(), "Failed"))
	}

	msg, err := h.au.ChangeEmail(*r)
	if err != nil {
		return c.JSON(RespondNotFound(err.Error(), "Failed"))
	}

	return c.JSON(RespondAccepted(msg, "Success"))
}
