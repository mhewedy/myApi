package commons

import "github.com/go-playground/validator"

var (
	Validate *validator.Validate
)

func InitValidate() {
	Validate = validator.New()
}
