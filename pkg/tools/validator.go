package tools

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate 驗證結構體
func Validate(data interface{}) error {
	return validate.Struct(data)
}

// ValidateVar 驗證單一變數
func ValidateVar(field interface{}, tag string) error {
	return validate.Var(field, tag)
}
