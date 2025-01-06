package api

import (
	"github.com/JueViGrace/bakery-go/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) UserRoutes(api fiber.Router) {
	usersGroup := api.Group("/users", a.sessionMiddleware)

	userHandler := handlers.NewUserHandler(a.db.UserStore())

	usersGroup.Get("/", a.sessionMiddleware, a.adminAuthMiddleware, userHandler.GetUsers)
	usersGroup.Get("/:id", a.sessionMiddleware, a.userIdMiddleware, userHandler.GetUserById)
	usersGroup.Patch("/:id", a.sessionMiddleware, a.userIdMiddleware, userHandler.UpdateUser)
	usersGroup.Delete("/:id", a.sessionMiddleware, a.userIdMiddleware, userHandler.DeleteUser)
}
