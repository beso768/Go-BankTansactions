package api

import (
	"github.com/beso768/Go-BankTansactions/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSuportedCurrency(currency)
	}
	return false
}
