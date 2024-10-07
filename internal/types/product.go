package types

import (
	"time"

	"github.com/JueViGrace/bakery-go/internal/database"
	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Price       string    `json:"price"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Stock       int32     `json:"stock"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"-"`
}

type CreateProductRequest struct {
	Price       string `json:"price"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Stock       int32  `json:"stock"`
	Image       string `json:"image"`
}

type UpdateProductRequest struct {
	Price       string    `json:"price"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Stock       int32     `json:"stock"`
	Image       string    `json:"image"`
	ID          uuid.UUID `json:"id"`
}

func DbProductToProduct(dp database.BakeryProduct) *Product {
	return &Product{
		ID:          dp.ID,
		Price:       dp.Price,
		Name:        dp.Name,
		Description: dp.Description,
		Category:    dp.Category,
		Stock:       dp.Stock,
		Image:       dp.Image,
		CreatedAt:   dp.CreatedAt.Time,
		UpdatedAt:   dp.UpdatedAt.Time,
		DeletedAt:   dp.DeletedAt.Time,
	}
}

func NewCreateProductParams(cr CreateProductRequest) (*database.CreateProductParams, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &database.CreateProductParams{
		ID:          id,
		Price:       cr.Price,
		Name:        cr.Name,
		Description: cr.Description,
		Category:    cr.Category,
		Stock:       cr.Stock,
		Image:       cr.Image,
	}, nil
}

func NewUpdateProductParams(ur UpdateProductRequest) *database.UpdateProductParams {
	return &database.UpdateProductParams{
		ID:          ur.ID,
		Price:       ur.Price,
		Name:        ur.Name,
		Description: ur.Description,
		Category:    ur.Category,
		Stock:       ur.Stock,
		Image:       ur.Image,
	}
}
