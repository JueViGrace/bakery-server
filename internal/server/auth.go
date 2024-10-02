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

type AuthHandler struct {
	as data.AuthStore
}

func NewAuthHandler(as data.AuthStore) AuthRoutes {
	return &AuthHandler{
		as: as,
	}
}

func (h *AuthHandler) SignIn(c *fiber.Ctx) error {
	r := new(data.SignInRequest)
	if err := c.BodyParser(r); err != nil {
		return RespondBadRequest(c, err.Error(), "Failed")
	}

	token, err := h.as.SignIn(*r)
	if err != nil {
		return RespondNotFound(c, err.Error(), "Failed")
	}

	return RespondOk(c, token, "Success")
}

func (h *AuthHandler) SignUp(c *fiber.Ctx) error {
	r := new(data.SignUpRequest)
	if err := c.BodyParser(r); err != nil {
		return RespondBadRequest(c, err.Error(), "Failed")
	}

	token, err := h.as.SignUp(*r)
	if err != nil {
		return RespondNotFound(c, err.Error(), "Failed")
	}

	return RespondCreated(c, token, "Success")
}

func (h *AuthHandler) RecoverPassword(c *fiber.Ctx) error {
	r := new(data.RecoverPasswordRequest)
	if err := c.BodyParser(r); err != nil {
		return RespondBadRequest(c, err.Error(), "Failed")
	}

	msg, err := h.as.RecoverPassword(*r)
	if err != nil {
		return RespondNotFound(c, err.Error(), "Failed")
	}

	return RespondAccepted(c, msg, "Success")
}

func (h *AuthHandler) ChangeEmail(c *fiber.Ctx) error {
	r := new(data.ChangeEmailRequest)
	if err := c.BodyParser(r); err != nil {
		return RespondBadRequest(c, err.Error(), "Failed")
	}

	msg, err := h.as.ChangeEmail(*r)
	if err != nil {
		return RespondNotFound(c, err.Error(), "Failed")
	}

	return RespondAccepted(c, msg, "Success")
}
