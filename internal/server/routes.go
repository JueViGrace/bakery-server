package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func (s *FiberServer) RegisterRoutes() {
	s.HealthRoute()

	s.UserRoutes()
	s.AuthRoutes()
	s.ProductRoutes()
}

func (s *FiberServer) HealthRoute() {
	s.App.Get("/api/health", func(c *fiber.Ctx) error {
		res := NewAPIResponse(http.StatusOK, http.StatusText(http.StatusOK), s.db.Health(), "Success")
		return c.Status(res.Status).JSON(res)
	})
}

func (s *FiberServer) MonitorRoute() {
	s.App.Get("/api/metrics", monitor.New())
}

func (s *FiberServer) protectedRoute() {}
