package api

import (
	"github.com/JueViGrace/bakery-server/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) AuthRoutes(api fiber.Router) {
	authGroup := api.Group("/auth")

	authHandler := handlers.NewAuthHandler(a.db.AuthStore())

	authGroup.Post("/signIn", authHandler.SignIn)
	authGroup.Post("/signUp", authHandler.SignUp)
	authGroup.Post("/refresh", authHandler.Refresh)
	authGroup.Post("/recover/password", authHandler.RecoverPassword)
}
