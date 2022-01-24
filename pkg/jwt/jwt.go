package jwt

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/ryanadiputraa/api-gervichstore.id/domain"
	"github.com/ryanadiputraa/api-gervichstore.id/pkg/wrapper"
)

func GenerateTokenWithClaims(claims *domain.Claims, secret []byte) (token string, err error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = at.SignedString(secret)
	if err != nil {
		return token, &wrapper.GenericError{
			HTTPCode: http.StatusInternalServerError,
			Code:     500,
			Message:  "Fail to generate token",
			Cause:    err.Error(),
		}
	}
	return
}

func ParseTokenWithClaims(tokenString string, claims *domain.Claims, secret []byte) (err error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return &wrapper.GenericError{
				HTTPCode: http.StatusUnauthorized,
				Code:     403,
				Message:  "Invalid signature",
				Cause:    err.Error(),
			}
		}
	}

	if !token.Valid {
		return &wrapper.GenericError{
			HTTPCode: http.StatusUnauthorized,
			Code:     403,
			Message:  "Invalid token",
		}
	}
	return
}

func ExtractTokenFromAuthorizationHeader(headers http.Header) (token string, err error) {
	header, ok := headers["Authorization"]
	if !ok {
		return token, &wrapper.GenericError{
			HTTPCode: http.StatusUnauthorized,
			Code:     403,
			Message:  "Unable to parse token",
			Cause:    "Missing authorization header",
		}
	}

	if header[0] == "" {
		return token, &wrapper.GenericError{
			HTTPCode: http.StatusUnauthorized,
			Code:     403,
			Message:  "Unable to parse token",
			Cause:    "Missing bearer token",
		}
	}

	headerValue := strings.Split(header[0], "")
	if len(headerValue) < 2 {
		return "", &wrapper.GenericError{
			HTTPCode: http.StatusUnauthorized,
			Code:     403,
			Message:  "Unable to parse token",
			Cause:    "Invalid token format",
		}
	}

	token = headerValue[1]
	return
}
