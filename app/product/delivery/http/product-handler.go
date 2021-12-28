package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ryanadiputraa/api-gervichstore.id/domain"
	"github.com/ryanadiputraa/api-gervichstore.id/pkg/wrapper"
)

type ProductHandler struct {
	productUsecase domain.IProductUsecase
}

func NewProductHandler(router, authRouter *mux.Router, productUsecase domain.IProductUsecase) {
	handler := &ProductHandler{productUsecase: productUsecase}

	router.HandleFunc("/api/products", handler.GetProducts).Methods(http.MethodGet)
	router.HandleFunc("/api/products/{id}", handler.GetProductByID).Methods(http.MethodGet)
}

func (h *ProductHandler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	products, err := h.productUsecase.GetProducts(ctx)
	if err != nil {
		if errVal, ok := err.(*wrapper.GenericError); ok {
			wrapper.WrapResponse(rw, errVal.HTTPCode, &wrapper.Response{
				Code:    errVal.Code,
				Message: errVal.Message,
				Error:   errVal.Cause,
			})
			return
		}
	}

	wrapper.WrapResponse(rw, http.StatusOK, &wrapper.Response{
		Code:    http.StatusOK,
		Message: wrapper.SuccessLabel,
		Data:    products,
	})
}

func (h *ProductHandler) GetProductByID(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	productID := mux.Vars(r)["id"]
	product, err := h.productUsecase.GetProductByID(ctx, productID)
	if err != nil {
		if errVal, ok := err.(*wrapper.GenericError); ok {
			wrapper.WrapResponse(rw, errVal.HTTPCode, &wrapper.Response{
				Code:    errVal.Code,
				Message: errVal.Message,
				Error:   errVal.Cause,
			})
			return
		}
	}

	wrapper.WrapResponse(rw, http.StatusOK, &wrapper.Response{
		Code:    http.StatusOK,
		Data:    product,
		Message: wrapper.SuccessLabel,
	})
}
