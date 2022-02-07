package usecase

import (
	"context"
	"net/http"
	"time"

	logging "github.com/sirupsen/logrus"

	"github.com/gocraft/dbr/v2"
	"github.com/ryanadiputraa/api-gervichstore.id/domain"
	"github.com/ryanadiputraa/api-gervichstore.id/pkg/timeparser"
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

func (u *ProductUsecase) AddProduct(ctx context.Context, payload domain.ProductDTO) (err error) {
	tx, err := u.writeSession.BeginTx(ctx, nil)
	if err != nil {
		logging.Error("Fail to begin db transactions: ", err.Error())
		return
	}
	defer tx.RollbackUnlessCommitted()

	payload.CreatedAt = timeparser.ConverTimeToProperFormat(time.Now())
	err = u.productRepository.Add(ctx, tx, payload)
	if err != nil {
		logging.Error("Fail to add product: ", err.Error())
		return
	}
	tx.Commit()
	return
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

func (u *ProductUsecase) UpdateProduct(ctx context.Context, productID string, payload domain.UpdateProductPayload) (err error) {
	tx, err := u.writeSession.BeginTx(ctx, nil)
	if err != nil {
		logging.Error("Fail to begin db transaction: ", err.Error())
		return
	}
	defer tx.RollbackUnlessCommitted()

	queryConditions := map[string]interface{}{
		"id": productID,
	}

	updatedTime := timeparser.ConverTimeToProperFormat(time.Now())
	updatePayload := map[string]interface{}{
		"product_name": payload.ProductName,
		"price":        payload.Price,
		"stock":        payload.Stock,
		"updated_at":   updatedTime,
	}

	err = u.productRepository.Update(ctx, tx, queryConditions, updatePayload)
	if err != nil {
		logging.Error("Fail to update product: ", err.Error())
		return
	}
	tx.Commit()
	return
}
