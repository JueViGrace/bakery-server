package api

import "github.com/JueViGrace/bakery-go/internal/handlers"

func (a *api) ProductRoutes() {
	productRoutes := a.App.Group("/api/products")

	productHandler := handlers.NewProductHandler(a.db.ProductStore())

	productRoutes.Get("/", productHandler.GetProducts)
	productRoutes.Get("/:id", productHandler.GetProductById)
	productRoutes.Post("/", a.adminAuthMiddleware, productHandler.CreateProduct)
	productRoutes.Patch("/", a.adminAuthMiddleware, productHandler.UpdateProduct)
	productRoutes.Delete("/:id", a.adminAuthMiddleware, productHandler.DeleteProduct)
}
