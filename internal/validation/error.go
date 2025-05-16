package validation

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Error struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Message  string `json:"message"`
}

func GetValidationErrors(request *http.Request, err error) *[]Error {
	var validationErrors []Error
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		trans := GetTranslatorFromHeader(request.Header.Get("Accept-Language"))

		for _, fe := range ve {
			validationErrors = append(validationErrors, Error{
				Property: fe.Field(),
				Tag:      fe.Tag(),
				Message:  fe.Translate(trans),
			})
		}
		return &validationErrors
	}
	return nil
}
