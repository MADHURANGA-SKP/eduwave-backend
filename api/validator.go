package api

import (
	"eduwave-back-end/util"

	"github.com/go-playground/validator/v10"
)

var validUsername validator.Func = func(fl validator.FieldLevel) bool {
	if username, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportedUsername(username) == nil
	}
	return false
}
