package handlers

import (
	"github.com/JueViGrace/bakery-server/internal/data"
	"github.com/JueViGrace/bakery-server/internal/types"
	"github.com/JueViGrace/bakery-server/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler interface {
	GetUsers(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx, userId *uuid.UUID) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type userHandler struct {
	db data.UserStore
}

func NewUserHandler(db data.UserStore) UserHandler {
	return &userHandler{
		db: db,
	}
}

func (h *userHandler) GetUsers(c *fiber.Ctx) (err error) {
	users, err := h.db.GetUsers()
	if err != nil {
		res := types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res := types.RespondOk(users, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *userHandler) GetUserById(c *fiber.Ctx, userId *uuid.UUID) error {
	user, err := h.db.GetUserById(userId)
	if err != nil {
		res := types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res := types.RespondOk(user, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	r := new(types.UpdateUserRequest)
	if err := c.BodyParser(r); err != nil {
		res := types.RespondBadRequest(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	user, err := h.db.UpdateUser(r)
	if err != nil {
		res := types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res := types.RespondAccepted(user, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *userHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := util.GetIdFromParams(c.Params("id"))
	if err != nil {
		res := types.RespondBadRequest(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	err = h.db.DeleteUser(id)
	if err != nil {
		res := types.RespondBadRequest(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res := types.RespondNoContent("Deleted", "Success")
	return c.Status(res.Status).JSON(res)
}
