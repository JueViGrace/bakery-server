package api

import (
	"github.com/JueViGrace/bakery-go/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) ProductRoutes(api fiber.Router) {
	productRoutes := api.Group("/api/products")

	productHandler := handlers.NewProductHandler(a.db.ProductStore())

	productRoutes.Get("/", productHandler.GetProducts)
	productRoutes.Get("/:id", productHandler.GetProductById)
	productRoutes.Post("/", a.sessionMiddleware, a.adminAuthMiddleware, productHandler.CreateProduct)
	productRoutes.Patch("/", a.sessionMiddleware, a.adminAuthMiddleware, productHandler.UpdateProduct)
	productRoutes.Delete("/:id", a.sessionMiddleware, a.adminAuthMiddleware, productHandler.DeleteProduct)
}
