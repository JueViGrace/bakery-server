package api

import (
	"github.com/JueViGrace/bakery-go/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func (a *api) RegisterRoutes() {
	a.HealthRoute()
	a.MonitorRoute()

	a.UserRoutes()
	a.AuthRoutes()
	a.ProductRoutes()
}

func (s *api) HealthRoute() {
	s.App.Get("/api/health", s.adminAuthMiddleware, func(c *fiber.Ctx) error {
		res := types.RespondOk(s.db.Health(), "Success")
		return c.Status(res.Status).JSON(res)
	})
}

func (s *api) MonitorRoute() {
	s.App.Get("/api/metrics", s.adminAuthMiddleware, monitor.New())
}
