package api

import (
	"time"

	"github.com/JueViGrace/bakery-go/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func (a *api) RegisterRoutes() {
	a.ApiRoutes()
	a.WebRoutes()
}

func (a *api) ApiRoutes() {
	api := a.App.Group("/api")

	api.Get("/health", a.HealthRoute)
	api.Get("/metrics", monitor.New(monitor.Config{
		Refresh: time.Duration(time.Second),
	}))

	a.UserRoutes(api)
	a.AuthRoutes(api)
	a.ProductRoutes(api)
}

func (a *api) WebRoutes() {
	a.App.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Root")
	})
}

func (s *api) HealthRoute(c *fiber.Ctx) error {
	res := types.RespondOk(s.db.Health(), "Success")
	return c.Status(res.Status).JSON(res)
}
