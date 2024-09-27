package server

import (
	"github.com/JueViGrace/bakery-go/internal/database"
	"github.com/gofiber/fiber/v2"
)

type UserRoutes interface {
	GetUsers(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type UserHandler struct {
	db database.Service
}

func NewUserHandler(db database.Service) UserRoutes {
	return &UserHandler{
		db: db,
	}
}

func (s *FiberServer) UserRoutes() {
	userGroup := s.App.Group("/api/users")

	userHandler := NewUserHandler(s.db)

	userGroup.Get("/", userHandler.GetUsers)
	userGroup.Get("/:id", userHandler.GetUserById)
	userGroup.Patch("/:id", userHandler.UpdateUser)
	userGroup.Delete("/:id", userHandler.DeleteUser)

}

func (h *UserHandler) GetUsers(c *fiber.Ctx) (err error) {
	res := RespondOk(map[string]string{"success": "Hello world!"}, "Success")
	return c.JSON(res)
}

func (h *UserHandler) GetUserById(c *fiber.Ctx) (err error) {

	return
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) (err error) {

	return
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) (err error) {

	return
}
