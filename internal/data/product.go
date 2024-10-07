package data

import (
	"context"

	"github.com/JueViGrace/bakery-go/internal/database"
	"github.com/JueViGrace/bakery-go/internal/types"
	"github.com/google/uuid"
)

type ProductStore interface {
	GetProducts() ([]*types.Product, error)
	GetProductById(id uuid.UUID) (*types.Product, error)
	CreateProduct(cr types.CreateProductRequest) (*types.Product, error)
	UpdateProduct(ur types.UpdateProductRequest) (*types.Product, error)
	DeleteProduct(id uuid.UUID) error
}

func (s *storage) ProductStore() ProductStore {
	return NewProductStore(s.ctx, s.queries)
}

type productStore struct {
	ctx context.Context
	db  *database.Queries
}

func NewProductStore(ctx context.Context, db *database.Queries) ProductStore {
	return &productStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *productStore) GetProducts() ([]*types.Product, error) {
	products := make([]*types.Product, 0)

	dbProducts, err := s.db.GetProducts(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, p := range dbProducts {
		products = append(products, types.DbProductToProduct(p))
	}

	return products, nil
}

func (s *productStore) GetProductById(id uuid.UUID) (*types.Product, error) {
	product := new(types.Product)

	dbProduct, err := s.db.GetProductById(s.ctx, id)
	if err != nil {
		return nil, err
	}

	product = types.DbProductToProduct(dbProduct)

	return product, nil
}

func (s *productStore) CreateProduct(cr types.CreateProductRequest) (*types.Product, error) {
	product := new(types.Product)

	r, err := types.NewCreateProductParams(cr)
	if err != nil {
		return nil, err
	}

	dbProduct, err := s.db.CreateProduct(s.ctx, *r)
	if err != nil {
		return nil, err
	}

	product = types.DbProductToProduct(dbProduct)

	return product, nil
}

func (s *productStore) UpdateProduct(ur types.UpdateProductRequest) (*types.Product, error) {
	product := new(types.Product)

	dbProduct, err := s.db.UpdateProduct(s.ctx, *types.NewUpdateProductParams(ur))
	if err != nil {
		return nil, err
	}

	product = types.DbProductToProduct(dbProduct)

	return product, nil
}

func (s *productStore) DeleteProduct(id uuid.UUID) error {
	err := s.db.DeleteProduct(s.ctx, id)
	if err != nil {
		return err
	}

	return nil
}
