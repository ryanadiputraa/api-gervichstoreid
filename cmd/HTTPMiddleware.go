package cmd

import (
	"net/http"
	"runtime/debug"

	logging "github.com/sirupsen/logrus"
)

func CorsMiddleware() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			rw.WriteHeader(http.StatusOK)
		}
	})
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				debug.PrintStack()
				logging.Error("Panic errpr: %s", err)
				return
			}
		}()
		next.ServeHTTP(rw, r)
	})
}
