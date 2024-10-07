package api

import "github.com/JueViGrace/bakery-go/internal/handlers"

func (a *api) UserRoutes() {
	usersGroup := a.App.Group("/api/users", a.authMiddleware)

	userHandler := handlers.NewUserHandler(a.db.UserStore())

	usersGroup.Get("/", a.adminAuthMiddleware, userHandler.GetUsers)
	usersGroup.Get("/:id", a.checkUserIdParamMiddleware, userHandler.GetUserById)
	usersGroup.Patch("/", a.checkUpdateUserMiddleware, userHandler.UpdateUser)
	usersGroup.Delete("/:id", a.checkUserIdParamMiddleware, userHandler.DeleteUser)
}
