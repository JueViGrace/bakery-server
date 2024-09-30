package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterRoutes() {
	s.HealthRoute()

	s.UserRoutes()
    s.AuthRoutes()
}

func (s *FiberServer) HealthRoute() {
	s.App.Get("/api/health", func(c *fiber.Ctx) error {
		res := NewAPIResponse(http.StatusOK, http.StatusText(http.StatusOK), s.db.Health(), "Success")
		return c.Status(res.Status).JSON(res)
	})
}
