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

type userHandler struct {
	us data.UserStore
}

func NewUserHandler(us data.UserStore) UserRoutes {
	return &userHandler{
		us: us,
	}
}

func (s *FiberServer) UserRoutes() {
	usersGroup := s.App.Group("/api/users")

	userHandler := NewUserHandler(s.db.UserStore())

	usersGroup.Get("/", userHandler.GetUsers)
	usersGroup.Get("/:id", userHandler.GetUserById)
	usersGroup.Patch("/", userHandler.UpdateUser)
	usersGroup.Delete("/:id", userHandler.DeleteUser)
}

func (h *userHandler) GetUsers(c *fiber.Ctx) (err error) {
	users, err := h.us.GetUsers()
	if err != nil {
		return c.JSON(RespondNotFound(err.Error(), "Failed"))
	}

	return c.JSON(RespondOk(users, "Success"))
}

func (h *userHandler) GetUserById(c *fiber.Ctx) error {
	id, err := util.GetIdFromParams(c)
	if err != nil {
		return c.JSON(RespondBadRequest(err.Error(), "Failed"))
	}

	user, err := h.us.GetUser(*id)
	if err != nil {
		return c.JSON(RespondNotFound(user, err.Error()))
	}

	return c.JSON(RespondOk(user, "Success"))
}

func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	ur := new(data.UpdateUserRequest)
	if err := c.BodyParser(ur); err != nil {
		return c.JSON(RespondBadRequest(err.Error(), "Failure"))
	}

	user, err := h.us.UpdateUser(*ur)
	if err != nil {
		return c.JSON(RespondNotFound(err.Error(), "Failure"))
	}

	return c.JSON(RespondAccepted(user, "Success"))
}

func (h *userHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := util.GetIdFromParams(c)
	if err != nil {
		return c.JSON(RespondBadRequest(err.Error(), "Failed"))
	}

	err = h.us.DeleteUser(*id)
	if err != nil {
		return c.JSON(RespondBadRequest(err.Error(), "Failed"))
	}

	return c.JSON(RespondNoContent("Deleted", "Success"))
}
