package server

import (
	"github.com/JueViGrace/bakery-go/internal/data"
	"github.com/JueViGrace/bakery-go/internal/util"
	"github.com/gofiber/fiber/v2"
)

type ProductRoutes interface {
	GetProducts(c *fiber.Ctx) error
	GetProduct(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
}

type productHandler struct {
	ps data.ProductStore
}

func NewProductHandler(ps data.ProductStore) ProductRoutes {
	return &productHandler{
		ps: ps,
	}
}

func (s *FiberServer) ProductRoutes() {
	productRoutes := s.Group("/api/products")

	productHandler := NewProductHandler(s.db.ProductStore())

	productRoutes.Get("/", productHandler.GetProducts)
	productRoutes.Get("/:id", productHandler.GetProduct)
	productRoutes.Post("/", productHandler.CreateProduct)
	productRoutes.Patch("/", productHandler.UpdateProduct)
	productRoutes.Delete("/:id", productHandler.DeleteProduct)
}

func (h *productHandler) GetProducts(c *fiber.Ctx) error {
	products, err := h.ps.GetProducts()
	if err != nil {
		return c.JSON(RespondNotFound(err.Error(), "Failed"))
	}

	return c.JSON(RespondOk(products, "Success"))
}

func (h *productHandler) GetProduct(c *fiber.Ctx) error {
	id, err := util.GetIdFromParams(c)
	if err != nil {
		return c.JSON(RespondBadRequest(err.Error(), "Failed"))
	}

	product, err := h.ps.GetProductById(*id)
	if err != nil {
		return c.JSON(RespondNotFound(err.Error(), "Failed"))
	}

	return c.JSON(RespondOk(product, "Success"))
}

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	cr := new(data.CreateProductRequest)
	if err := c.BodyParser(cr); err != nil {
		return c.JSON(RespondBadRequest(err.Error(), "Failed"))
	}

	product, err := h.ps.CreateProduct(*cr)
	if err != nil {
		return c.JSON(RespondNotFound(err.Error(), "Failed"))
	}

	return c.JSON(RespondCreated(product, "Success"))
}

func (h *productHandler) UpdateProduct(c *fiber.Ctx) error {
	ur := new(data.UpdateProductRequest)
	if err := c.BodyParser(ur); err != nil {
		return c.JSON(RespondBadRequest(err.Error(), "Failed"))
	}

	product, err := h.ps.UpdateProduct(*ur)
	if err != nil {
		return c.JSON(RespondNoContent(err.Error(), "Failed"))
	}

	return c.JSON(RespondAccepted(product, "Success"))
}

func (h *productHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := util.GetIdFromParams(c)
	if err != nil {
		return c.JSON(RespondBadRequest(err.Error(), "Failed"))
	}

	err = h.ps.DeleteProduct(*id)
	if err != nil {
		return c.JSON(RespondNoContent(err.Error(), "Failed"))
	}

	return c.JSON(RespondNoContent("Deleted!", "Success"))
}
