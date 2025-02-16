package api

import (
	"github.com/JueViGrace/bakery-server/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) AuthRoutes(api fiber.Router) {
	authGroup := api.Group("/auth")

	authHandler := handlers.NewAuthHandler(a.db.AuthStore(), a.validator)

	authGroup.Post("/signIn", authHandler.SignIn)
	authGroup.Post("/signUp", authHandler.SignUp)
	authGroup.Post("/refresh", a.authenticatedHandler(authHandler.Refresh))
	authGroup.Post("/recover/password", authHandler.RecoverPassword)
}
