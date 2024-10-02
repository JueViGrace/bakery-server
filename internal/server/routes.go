package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func (s *FiberServer) RegisterRoutes() {
	s.HealthRoute()
	s.MonitorRoute()

	s.UserRoutes()
	s.AuthRoutes()
	s.ProductRoutes()
}

func (s *FiberServer) HealthRoute() {
	s.App.Get("/api/health", s.adminAuthMiddleware, func(c *fiber.Ctx) error {
		return RespondOk(c, s.db.Health(), "Success")
	})
}

func (s *FiberServer) MonitorRoute() {
	s.App.Get("/api/metrics", s.adminAuthMiddleware, monitor.New())
}

func (s *FiberServer) AuthRoutes() {
	authGroup := s.Group("/api/auth")

	authHandler := NewAuthHandler(s.db.AuthStore())

	authGroup.Post("/signIn", authHandler.SignIn)
	authGroup.Post("/signUp", authHandler.SignUp)
	authGroup.Post("/recover/password", authHandler.RecoverPassword)
	authGroup.Post("/recover/email", authHandler.ChangeEmail)
}

func (s *FiberServer) UserRoutes() {
	usersGroup := s.App.Group("/api/users", s.authMiddleware)

	userHandler := NewUserHandler(s.db.UserStore())

	usersGroup.Get("/", s.adminAuthMiddleware, userHandler.GetUsers)
	usersGroup.Get("/:id", s.checkUserIdParamMiddleware, userHandler.GetUserById)
	usersGroup.Patch("/", s.checkUpdateUserMiddleware, userHandler.UpdateUser)
	usersGroup.Delete("/:id", s.checkUserIdParamMiddleware, userHandler.DeleteUser)
}

func (s *FiberServer) ProductRoutes() {
	productRoutes := s.Group("/api/products")

	productHandler := NewProductHandler(s.db.ProductStore())

	productRoutes.Get("/", productHandler.GetProducts)
	productRoutes.Get("/:id", productHandler.GetProductById)
	productRoutes.Post("/", s.adminAuthMiddleware, productHandler.CreateProduct)
	productRoutes.Patch("/", s.adminAuthMiddleware, productHandler.UpdateProduct)
	productRoutes.Delete("/:id", s.adminAuthMiddleware, productHandler.DeleteProduct)
}
