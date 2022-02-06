package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ryanadiputraa/api-gervichstore.id/domain"
	"github.com/ryanadiputraa/api-gervichstore.id/pkg/wrapper"
)

type AuthHandler struct {
	authUsecase domain.IAuthUsecase
}

func NewAuthHandler(router *mux.Router, authUsecase domain.IAuthUsecase) {
	handler := &AuthHandler{authUsecase: authUsecase}

	router.HandleFunc("/register", handler.Register).Methods(http.MethodPost)
}

func (h *AuthHandler) Register(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload domain.UserDTO

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

	err = h.authUsecase.Register(ctx, payload)
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
