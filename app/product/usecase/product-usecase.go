package usecase

import (
	"context"
	"net/http"

	logging "github.com/sirupsen/logrus"

	"github.com/gocraft/dbr/v2"
	"github.com/ryanadiputraa/api-gervichstore.id/domain"
	"github.com/ryanadiputraa/api-gervichstore.id/pkg/wrapper"
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

	if productS == nil {
		return product, &wrapper.GenericError{
			HTTPCode: http.StatusNotFound,
			Code:     404,
			Message:  wrapper.BadRequestLabel,
		}
	}

	product = *productS
	return
}
