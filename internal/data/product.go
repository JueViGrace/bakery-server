package data

import (
	"context"
	"time"

	"github.com/JueViGrace/bakery-go/internal/db"
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

type ProductStore interface {
	GetProducts() ([]*Product, error)
	GetProductById(id uuid.UUID) (*Product, error)
	CreateProduct(cr CreateProductRequest) (*Product, error)
	UpdateProduct(ur UpdateProductRequest) (*Product, error)
	DeleteProduct(id uuid.UUID) error
}

type productStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewProductStore(ctx context.Context, db *db.Queries) ProductStore {
	return &productStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *productStore) GetProducts() ([]*Product, error) {
	products := make([]*Product, 0)

	dbProducts, err := s.db.GetProducts(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, p := range dbProducts {
		products = append(products, dbProductToProduct(p))
	}

	return products, nil
}

func (s *productStore) GetProductById(id uuid.UUID) (*Product, error) {
	product := new(Product)

	dbProduct, err := s.db.GetProductById(s.ctx, id)
	if err != nil {
		return nil, err
	}

	product = dbProductToProduct(dbProduct)

	return product, nil
}

func (s *productStore) CreateProduct(cr CreateProductRequest) (*Product, error) {
	product := new(Product)

	r, err := newCreateProductParams(cr)
	if err != nil {
		return nil, err
	}

	dbProduct, err := s.db.CreateProduct(s.ctx, *r)
	if err != nil {
		return nil, err
	}

	product = dbProductToProduct(dbProduct)

	return product, nil
}

func (s *productStore) UpdateProduct(ur UpdateProductRequest) (*Product, error) {
	product := new(Product)

	dbProduct, err := s.db.UpdateProduct(s.ctx, *newUpdateProductParams(ur))
	if err != nil {
		return nil, err
	}

	product = dbProductToProduct(dbProduct)

	return product, nil
}

func (s *productStore) DeleteProduct(id uuid.UUID) error {
	err := s.db.DeleteProduct(s.ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func dbProductToProduct(dp db.BakeryProduct) *Product {
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

func newCreateProductParams(cr CreateProductRequest) (*db.CreateProductParams, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &db.CreateProductParams{
		ID:          id,
		Price:       cr.Price,
		Name:        cr.Name,
		Description: cr.Description,
		Category:    cr.Category,
		Stock:       cr.Stock,
		Image:       cr.Image,
	}, nil
}

func newUpdateProductParams(ur UpdateProductRequest) *db.UpdateProductParams {
	return &db.UpdateProductParams{
		ID:          ur.ID,
		Price:       ur.Price,
		Name:        ur.Name,
		Description: ur.Description,
		Category:    ur.Category,
		Stock:       ur.Stock,
		Image:       ur.Image,
	}
}
