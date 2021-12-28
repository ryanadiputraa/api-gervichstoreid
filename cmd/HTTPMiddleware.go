package cmd

import (
	"net/http"
	"runtime/debug"

	"github.com/ryanadiputraa/api-gervichstore.id/domain"
	JWTHelper "github.com/ryanadiputraa/api-gervichstore.id/pkg/jwt"
	"github.com/ryanadiputraa/api-gervichstore.id/pkg/wrapper"
	logging "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func CorsMiddleware() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			rw.WriteHeader(http.StatusOK)
		}
	})
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		tokenString, err := JWTHelper.ExtractTokenFromAuthorizationHeader(r.Header)
		if err != nil {
			if errVal, ok := err.(*wrapper.GenericError); ok {
				wrapper.WrapResponse(rw, errVal.HTTPCode, &wrapper.Response{
					Code:    errVal.Code,
					Message: errVal.Message,
					Error:   errVal.Cause,
				})
				return
			}
			return
		}

		claims := &domain.Claims{}
		secret := []byte(viper.GetString("jwt.secret.access"))
		err = JWTHelper.ParseTokenWithClaims(tokenString, claims, secret)
		if err != nil {
			if errVal, ok := err.(*wrapper.GenericError); ok {
				wrapper.WrapResponse(rw, errVal.HTTPCode, &wrapper.Response{
					Code:    errVal.Code,
					Message: errVal.Message,
					Error:   errVal.Cause,
				})
				return
			}
			return
		}

		r = r.WithContext(domain.WithUserClaims(r.Context(), claims))
		next.ServeHTTP(rw, r)
	})
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				debug.PrintStack()
				logging.Error("Panic error: ", err)
				return
			}
		}()
		next.ServeHTTP(rw, r)
	})
}
