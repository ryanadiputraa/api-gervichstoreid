package http

import (
	"encoding/json"
	"io/ioutil"
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

	router.HandleFunc("/products", handler.AddProduct).Methods(http.MethodPost)
	router.HandleFunc("/products", handler.GetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products/{id}", handler.GetProductByID).Methods(http.MethodGet)
	router.HandleFunc("/products/{id}", handler.UpdateProduct).Methods(http.MethodPut)
}

func (h *ProductHandler) AddProduct(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload domain.ProductDTO

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wrapper.WrapResponse(rw, http.StatusBadRequest, &wrapper.Response{
			Code:    http.StatusBadRequest,
			Message: wrapper.BadRequestLabel,
			Error:   err.Error(),
		})
		return
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		wrapper.WrapResponse(rw, http.StatusBadRequest, &wrapper.Response{
			Code:    http.StatusBadRequest,
			Message: wrapper.BadRequestLabel,
			Error:   err.Error(),
		})
		return
	}

	err = h.productUsecase.AddProduct(ctx, payload)
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

	wrapper.WrapResponse(rw, http.StatusCreated, &wrapper.Response{
		Code:    http.StatusCreated,
		Message: wrapper.SuccessLabel,
		Error:   "",
	})
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

func (h *ProductHandler) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var updatePayload domain.UpdateProductPayload

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wrapper.WrapResponse(rw, http.StatusBadRequest, &wrapper.Response{
			Code:    http.StatusBadRequest,
			Message: wrapper.BadRequestLabel,
			Error:   err.Error(),
		})
		return
	}

	err = json.Unmarshal(body, &updatePayload)
	if err != nil {
		wrapper.WrapResponse(rw, http.StatusBadRequest, &wrapper.Response{
			Code:    http.StatusBadRequest,
			Message: wrapper.BadRequestLabel,
			Error:   err.Error(),
		})
		return
	}

	productID := mux.Vars(r)["id"]
	err = h.productUsecase.UpdateProduct(ctx, productID, updatePayload)
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
		Error:   "",
	})
}
