package validators

import (
	"github.com/go-playground/validator"
	"github.com/kengo-k/password-manager/model"
)

func ValidateCategory(cmap map[string]*model.Category) func(f validator.FieldLevel) bool {
	return func(f validator.FieldLevel) bool {
		value := f.Field().String()
		_, ok := cmap[value]
		return ok
	}
}
