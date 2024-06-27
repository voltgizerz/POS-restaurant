package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetFirstValidatorError(valErr error) (err error) {
	for _, valErr := range valErr.(validator.ValidationErrors) {
		err = fmt.Errorf("field %s %s", valErr.Field(), valErr.Tag())
		continue
	}

	return err
}
