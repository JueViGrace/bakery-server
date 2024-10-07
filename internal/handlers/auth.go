package handlers

import (
	"github.com/JueViGrace/bakery-go/internal/data"
	"github.com/JueViGrace/bakery-go/internal/types"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	SignIn(c *fiber.Ctx) error
	SignUp(c *fiber.Ctx) error
	RecoverPassword(c *fiber.Ctx) error
	ChangeEmail(c *fiber.Ctx) error
}

type authHandler struct {
	db data.AuthStore
}

func NewAuthHandler(db data.AuthStore) AuthHandler {
	return &authHandler{
		db: db,
	}
}

func (h *authHandler) SignIn(c *fiber.Ctx) error {
	r := new(types.SignInRequest)
	if err := c.BodyParser(r); err != nil {
		res := types.RespondBadRequest(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	token, err := h.db.SignIn(*r)
	if err != nil {
		res := types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res := types.RespondOk(token, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) SignUp(c *fiber.Ctx) error {
	r := new(types.SignUpRequest)
	if err := c.BodyParser(r); err != nil {
		res := types.RespondBadRequest(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	token, err := h.db.SignUp(*r)
	if err != nil {
		res := types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res := types.RespondCreated(token, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) RecoverPassword(c *fiber.Ctx) error {
	r := new(types.RecoverPasswordRequest)
	if err := c.BodyParser(r); err != nil {
		res := types.RespondBadRequest(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	msg, err := h.db.RecoverPassword(*r)
	if err != nil {
		res := types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res := types.RespondAccepted(msg, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) ChangeEmail(c *fiber.Ctx) error {
	r := new(types.ChangeEmailRequest)
	if err := c.BodyParser(r); err != nil {
		res := types.RespondBadRequest(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	msg, err := h.db.ChangeEmail(*r)
	if err != nil {
		res := types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res := types.RespondAccepted(msg, "Success")
	return c.Status(res.Status).JSON(res)
}
