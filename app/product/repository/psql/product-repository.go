package psql

import (
	"context"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gocraft/dbr/v2"
	"github.com/ryanadiputraa/api-gervichstore.id/domain"
	"github.com/ryanadiputraa/api-gervichstore.id/pkg/wrapper"
)

type ProductRepository struct {
	sessionRead  *dbr.Session
	sessionWrite *dbr.Session
}

func NewProductRepository(read, write *dbr.Session) domain.IProductRepository {
	return &ProductRepository{
		sessionRead:  read,
		sessionWrite: write,
	}
}

func (r *ProductRepository) Add(ctx context.Context, tx *dbr.Tx, payload domain.ProductDTO) (err error) {
	tagList := wrapper.GetStructTagList(domain.ProductDTO{}, "db")
	errHystrix := hystrix.Do("SimpleQuery", func() error {
		_, err = tx.InsertInto("products").Columns(tagList...).Record(payload).ExecContext(ctx)
		return err
	}, nil)

	if errHystrix != nil {
		return &wrapper.GenericError{
			HTTPCode: http.StatusInternalServerError,
			Code:     500,
			Message:  wrapper.InternalServerErrorLabel,
			Cause:    errHystrix.Error(),
		}
	}
	return
}

func (r *ProductRepository) Fetch(ctx context.Context, readSession *dbr.Session, conditions map[string]interface{}) (products []domain.Product, err error) {
	tagList := wrapper.GetStructTagList(domain.Product{}, "db")
	errHystrix := hystrix.Do("SimpleQuery", func() error {
		db := readSession.Select(tagList...).From("products")

		for k, v := range conditions {
			db.Where(dbr.Eq(k, v))
		}

		_, err := db.LoadContext(ctx, &products)
		return err
	}, nil)

	if errHystrix != nil {
		return products, &wrapper.GenericError{
			HTTPCode: http.StatusInternalServerError,
			Code:     500,
			Message:  wrapper.InternalServerErrorLabel,
			Cause:    errHystrix.Error(),
		}
	}

	if products == nil {
		products = make([]domain.Product, 0)
	}

	return
}

func (r *ProductRepository) Query(ctx context.Context, readSession *dbr.Session, conditions map[string]interface{}) (product *domain.Product, err error) {
	tagList := wrapper.GetStructTagList(domain.Product{}, "db")
	errHystrix := hystrix.Do("SimpleQuery", func() error {
		db := readSession.Select(tagList...).From("products")

		for k, v := range conditions {
			db.Where(dbr.Eq(k, v))
		}

		_, err := db.LoadContext(ctx, &product)

		return err
	}, nil)

	if errHystrix != nil {
		return product, &wrapper.GenericError{
			HTTPCode: http.StatusInternalServerError,
			Code:     500,
			Message:  wrapper.InternalServerErrorLabel,
			Cause:    errHystrix.Error(),
		}
	}

	return
}

func (r *ProductRepository) Update(ctx context.Context, tx *dbr.Tx, conditions map[string]interface{}, payload map[string]interface{}) (err error) {
	errHystrix := hystrix.Do("SimpleQuery", func() error {
		db := tx.Update("products")

		for k, v := range conditions {
			db.Where(dbr.Eq(k, v))
		}

		_, err := db.SetMap(payload).ExecContext(ctx)
		return err
	}, nil)

	if errHystrix != nil {
		return &wrapper.GenericError{
			HTTPCode: http.StatusInternalServerError,
			Code:     500,
			Message:  wrapper.InternalServerErrorLabel,
			Cause:    err.Error(),
		}
	}
	return
}
