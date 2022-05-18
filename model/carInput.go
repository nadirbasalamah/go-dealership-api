package model

import "github.com/go-playground/validator/v10"

// CarInput represents input for car data
type CarInput struct {
	Brand string `validate:"required"`
	Name  string `validate:"required"`
	Year  int    `validate:"required"`
	Price int    `validate:"required"`
}

func (carInput CarInput) ValidateStruct() []*ErrorResponse {
	var errors []*ErrorResponse
	var validate *validator.Validate = validator.New()
	err := validate.Struct(carInput)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
	}

	return errors
}
