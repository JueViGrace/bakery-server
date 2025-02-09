package handlers

import (
	"fmt"
	"strings"

	"github.com/JueViGrace/bakery-server/internal/data"
	"github.com/JueViGrace/bakery-server/internal/types"
	"github.com/JueViGrace/bakery-server/internal/util"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	SignIn(c *fiber.Ctx) error
	SignUp(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx, a *types.AuthData) error
	RecoverPassword(c *fiber.Ctx) error
}

type authHandler struct {
	db        data.AuthStore
	validator *util.XValidator
}

func NewAuthHandler(db data.AuthStore, validator *util.XValidator) AuthHandler {
	return &authHandler{
		db:        db,
		validator: validator,
	}
}

func (h *authHandler) SignIn(c *fiber.Ctx) error {
	r := new(types.SignInRequest)
	res := new(types.APIResponse)

	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	if errs := h.validator.Validate(r); len(errs) > 0 {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		res = types.RespondBadRequest(nil, strings.Join(errMsgs, " and "))
		return c.Status(res.Status).JSON(res)
	}

	token, err := h.db.SignIn(r)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(token, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) SignUp(c *fiber.Ctx) error {
	r := new(types.SignUpRequest)
	res := new(types.APIResponse)

	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	if errs := h.validator.Validate(r); len(errs) > 0 {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		res = types.RespondBadRequest(nil, strings.Join(errMsgs, " and "))
		return c.Status(res.Status).JSON(res)
	}

	token, err := h.db.SignUp(r)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondCreated(token, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) Refresh(c *fiber.Ctx, a *types.AuthData) error {
	r := new(types.RefreshRequest)
	res := new(types.APIResponse)

	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	if errs := h.validator.Validate(r); len(errs) > 0 {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		res = types.RespondBadRequest(nil, strings.Join(errMsgs, " and "))
		return c.Status(res.Status).JSON(res)
	}

	msg, err := h.db.Refresh(r, a)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondAccepted(msg, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) RecoverPassword(c *fiber.Ctx) error {
	r := new(types.RecoverPasswordRequest)
	res := new(types.APIResponse)

	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	if errs := h.validator.Validate(r); len(errs) > 0 {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		res = types.RespondBadRequest(nil, strings.Join(errMsgs, " and "))
		return c.Status(res.Status).JSON(res)
	}

	msg, err := h.db.RecoverPassword(r)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondAccepted(msg, "Success")
	return c.Status(res.Status).JSON(res)
}
