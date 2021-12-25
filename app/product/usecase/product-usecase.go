package usecase

import (
	"context"

	logging "github.com/sirupsen/logrus"

	"github.com/gocraft/dbr/v2"
	"gitlab.com/ryanadiputraa/api-gervichstore.id/domain"
)

type ProductUsecase struct {
	readSession       *dbr.Session
	writeSession      *dbr.Session
	productRepository domain.IProductRepository
}

func NewProductUseCase(read, write *dbr.Session, productRepository domain.IProductRepository) domain.IProductUsecase {
	return &ProductUsecase{
		readSession:       read,
		writeSession:      write,
		productRepository: productRepository,
	}
}

func (p *ProductUsecase) GetProducts(ctx context.Context) (products []domain.Product, err error) {
	fetchConditions := make(map[string]interface{})
	products, err = p.productRepository.Fetch(ctx, p.readSession, fetchConditions)
	if err != nil {
		logging.Error("Fail to get all products: %s", err.Error())
		return
	}
	return
}

func (p *ProductUsecase) GetProductByID(ctx context.Context, productID int) (product domain.Product, err error) {
	fetchConditions := map[string]interface{}{
		"id": productID,
	}
	productS, err := p.productRepository.Query(ctx, p.readSession, fetchConditions)
	if err != nil {
		logging.Error("Fail to get product by id: %s", err.Error())
		return
	}

	if productS == nil {
		logging.Error("product with ID of '%s' not found", productID)
		return
	}

	product = *productS
	return
}
