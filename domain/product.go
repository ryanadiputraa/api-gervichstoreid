package domain

import (
	"context"

	"github.com/gocraft/dbr/v2"
	"gopkg.in/guregu/null.v3"
)

type IProductUsecase interface {
	AddProduct(ctx context.Context, payload ProductDTO) error
	GetProducts(ctx context.Context) ([]Product, error)
	GetProductByID(ctx context.Context, productID string) (Product, error)
	UpdateProduct(ctx context.Context, productID string, payload UpdateProductPayload) error
	// DeleteProduct(ctx context.Context, productID string) (ProductDTO, error)
}

type IProductRepository interface {
	Add(ctx context.Context, tx *dbr.Tx, payload ProductDTO) error
	Fetch(ctx context.Context, readSession *dbr.Session, conditions map[string]interface{}) ([]Product, error)
	Query(ctx context.Context, readSession *dbr.Session, conditions map[string]interface{}) (*Product, error)
	Update(ctx context.Context, tx *dbr.Tx, conditions map[string]interface{}, payload map[string]interface{}) error
	// Delete(ctx context.Context, writeSession *dbr.Session) error
}

type Product struct {
	ID          int         `json:"id" db:"id"`
	ProductName string      `json:"product_name" db:"product_name"`
	Price       int         `json:"price" db:"price"`
	Stock       int         `json:"stock" db:"stock"`
	CreatedAt   string      `json:"created_at" db:"created_at"`
	UpdateAt    null.String `json:"updated_at" db:"updated_at"`
}

type ProductDTO struct {
	ProductName string      `json:"product_name" db:"product_name"`
	Price       int         `json:"price" db:"price"`
	Stock       int         `json:"stock" db:"stock"`
	CreatedAt   string      `json:"created_at" db:"created_at"`
	UpdateAt    null.String `json:"updated_at" db:"updated_at"`
}

type UpdateProductPayload struct {
	ProductName string `json:"product_name" db:"product_name"`
	Price       int    `json:"price" db:"price"`
	Stock       int    `json:"stock" db:"stock"`
}
