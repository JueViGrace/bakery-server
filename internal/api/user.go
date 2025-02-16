package api

import (
	"github.com/JueViGrace/bakery-server/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) UserRoutes(api fiber.Router) {
	usersGroup := api.Group("/users", a.sessionMiddleware)

	userHandler := handlers.NewUserHandler(a.db.UserStore(), a.validator)

	usersGroup.Get("/", a.adminAuthMiddleware, userHandler.GetUsers)
	usersGroup.Get("/me", a.authenticatedHandler(userHandler.GetUserById))
	usersGroup.Patch("/:id", a.authenticatedHandler(userHandler.UpdateUser))
	usersGroup.Delete("/:id", a.authenticatedHandler(userHandler.DeleteUser))
}
