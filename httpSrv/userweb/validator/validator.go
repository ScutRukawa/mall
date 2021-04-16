package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidatorMobile(f1 validator.FieldLevel) bool {
	mobile := f1.Field().String()

	pattern := `^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`
	ok, _ := regexp.MatchString(pattern, mobile)
	if !ok {
		return false
	}
	return true
}
