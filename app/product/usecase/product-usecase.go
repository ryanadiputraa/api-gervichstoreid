package usecase

import (
	"context"

	logging "github.com/sirupsen/logrus"

	"github.com/gocraft/dbr/v2"
	"github.com/ryanadiputraa/api-gervichstore.id/domain"
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

func (u *ProductUsecase) GetProducts(ctx context.Context) (products []domain.Product, err error) {
	fetchConditions := make(map[string]interface{})
	products, err = u.productRepository.Fetch(ctx, u.readSession, fetchConditions)
	if err != nil {
		logging.Error("Fail to get all products: ", err.Error())
		return
	}
	return
}

func (u *ProductUsecase) GetProductByID(ctx context.Context, productID string) (product domain.Product, err error) {
	fetchConditions := map[string]interface{}{
		"id": productID,
	}
	productS, err := u.productRepository.Query(ctx, u.readSession, fetchConditions)
	if err != nil {
		logging.Error("Fail to get product by id: ", productID)
		return
	}

	product = *productS
	return
}
