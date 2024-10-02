package server

import (
	"github.com/JueViGrace/bakery-go/internal/data"
	"github.com/JueViGrace/bakery-go/internal/util"
	"github.com/gofiber/fiber/v2"
)

type ProductRoutes interface {
	GetProducts(c *fiber.Ctx) error
	GetProductById(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
}

type ProductHandler struct {
	ps data.ProductStore
}

func NewProductHandler(ps data.ProductStore) ProductRoutes {
	return &ProductHandler{
		ps: ps,
	}
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	products, err := h.ps.GetProducts()
	if err != nil {
		return RespondNotFound(c, err.Error(), "Failed")
	}

	return RespondOk(c, products, "Success")
}

func (h *ProductHandler) GetProductById(c *fiber.Ctx) error {
	id, err := util.GetIdFromParams(c.Params("id"))
	if err != nil {
		return RespondBadRequest(c, err.Error(), "Failed")
	}

	product, err := h.ps.GetProductById(*id)
	if err != nil {
		return RespondNotFound(c, err.Error(), "Failed")
	}

	return RespondOk(c, product, "Success")
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	cr := new(data.CreateProductRequest)
	if err := c.BodyParser(cr); err != nil {
		return RespondBadRequest(c, err.Error(), "Failed")
	}

	product, err := h.ps.CreateProduct(*cr)
	if err != nil {
		return RespondNotFound(c, err.Error(), "Failed")
	}

	return RespondCreated(c, product, "Success")
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	ur := new(data.UpdateProductRequest)
	if err := c.BodyParser(ur); err != nil {
		return RespondBadRequest(c, err.Error(), "Failed")
	}

	product, err := h.ps.UpdateProduct(*ur)
	if err != nil {
		return RespondNoContent(c, err.Error(), "Failed")
	}

	return RespondAccepted(c, product, "Success")
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := util.GetIdFromParams(c.Params("id"))
	if err != nil {
		return RespondBadRequest(c, err.Error(), "Failed")
	}

	err = h.ps.DeleteProduct(*id)
	if err != nil {
		return RespondNoContent(c, err.Error(), "Failed")
	}

	return RespondNoContent(c, "Deleted!", "Success")
}
