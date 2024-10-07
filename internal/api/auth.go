package api

import "github.com/JueViGrace/bakery-go/internal/handlers"

func (a *api) AuthRoutes() {
	authGroup := a.App.Group("/api/auth")

	authHandler := handlers.NewAuthHandler(a.db.AuthStore())

	authGroup.Post("/signIn", authHandler.SignIn)
	authGroup.Post("/signUp", authHandler.SignUp)
	authGroup.Post("/recover/password", authHandler.RecoverPassword)
	authGroup.Post("/recover/email", authHandler.ChangeEmail)
}
