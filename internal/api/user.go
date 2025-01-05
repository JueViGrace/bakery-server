package api

import "github.com/JueViGrace/bakery-go/internal/handlers"

func (a *api) UserRoutes() {
	usersGroup := a.App.Group("/api/users", a.authMiddleware)

	userHandler := handlers.NewUserHandler(a.db.UserStore())

	usersGroup.Get("/", userHandler.GetUsers)
	usersGroup.Get("/:id", userHandler.GetUserById)
	usersGroup.Patch("/", userHandler.UpdateUser)
	usersGroup.Delete("/:id", userHandler.DeleteUser)
}
