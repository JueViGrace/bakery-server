package server

import (
	"github.com/JueViGrace/bakery-go/internal/data"
	"github.com/JueViGrace/bakery-go/internal/util"
	"github.com/gofiber/fiber/v2"
)

type UserRoutes interface {
	GetUsers(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type UserHandler struct {
	us data.UserStore
}

func NewUserHandler(us data.UserStore) UserRoutes {
	return &UserHandler{
		us: us,
	}
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) (err error) {
	users, err := h.us.GetUsers()
	if err != nil {
		return RespondNotFound(c, err.Error(), "Failed")
	}

	return RespondOk(c, users, "Success")
}

func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	id, err := util.GetIdFromParams(c.Params("id"))
	if err != nil {
		return RespondBadRequest(c, err.Error(), "Failed")
	}

	user, err := h.us.GetUserById(*id)
	if err != nil {
		return RespondNotFound(c, user, err.Error())
	}

	return RespondOk(c, user, "Success")
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	ur := new(data.UpdateUserRequest)
	if err := c.BodyParser(ur); err != nil {
		return RespondBadRequest(c, err.Error(), "Failure")
	}

	user, err := h.us.UpdateUser(*ur)
	if err != nil {
		return RespondNotFound(c, err.Error(), "Failure")
	}

	return RespondAccepted(c, user, "Success")
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := util.GetIdFromParams(c.Params("id"))
	if err != nil {
		return RespondBadRequest(c, err.Error(), "Failed")
	}

	err = h.us.DeleteUser(*id)
	if err != nil {
		return RespondBadRequest(c, err.Error(), "Failed")
	}

	return RespondNoContent(c, "Deleted", "Success")
}
