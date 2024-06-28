package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetFirstValidatorError(valErr error) (err error) {
	valErrs := valErr.(validator.ValidationErrors)
	if len(valErrs) != 0 {
		err = fmt.Errorf("field %s %s", valErrs[0].Field(), valErrs[0].Tag())
	}

	return err
}
