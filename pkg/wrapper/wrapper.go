package wrapper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

var (
	SuccessLabel             = "Success"
	BadRequestLabel          = "Bad Request"
	InternalServerErrorLabel = "Internal Server Error"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func WrapResponse(rw http.ResponseWriter, HTTPStatus int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(HTTPStatus)
	json.NewEncoder(rw).Encode(data)
}

func GetStructTagList(data interface{}, tag string) (fields []string) {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		panic("passed interface params must be struct or pointer to a struct")
	}

	for i := 0; i < val.NumField(); i++ {
		_type := val.Type().Field(i)

		if _tag := _type.Tag.Get(tag); _tag != "" && _tag != "-" {
			var commaIndex int
			if commaIndex = strings.Index(_tag, ","); commaIndex < 0 {
				commaIndex = len(_tag)
			}
			fieldName := _tag[:commaIndex]
			fields = append(fields, fieldName)
		}
	}
	return
}

type GenericError struct {
	HTTPCode int
	Code     int
	Message  string
	Cause    string
}

func (e *GenericError) Error() string {
	err := fmt.Sprintf("%+v", e.Cause)
	return err
}
