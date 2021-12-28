package domain

import (
	"context"

	"github.com/gocraft/dbr/v2"
	"gopkg.in/guregu/null.v3"
)

type IProductUsecase interface {
	GetProducts(ctx context.Context) ([]Product, error)
	GetProductByID(ctx context.Context, productID string) (Product, error)
}

type IProductRepository interface {
	Fetch(ctx context.Context, readSession *dbr.Session, conditions map[string]interface{}) ([]Product, error)
	Query(ctx context.Context, readSession *dbr.Session, conditions map[string]interface{}) (*Product, error)
}

type Product struct {
	ID          int         `json:"id" db:"id"`
	ProductName string      `json:"product_name" db:"product_name"`
	Price       int         `json:"price" db:"price"`
	Stock       int         `json:"stock" db:"stock"`
	CreatedAt   string      `json:"created_at" db:"created_at"`
	UpdateAt    null.String `json:"updated_at" db:"updated_at"`
}
