package cvalidator

import (
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func ConfigureValidator() {
	Validate = validator.New()
}

type Validatable interface {
	StructLevelValidation(sl validator.StructLevel)
}

func ValidateStruct(validatable Validatable) error {
	result := validator.New()

	err := result.Struct(validatable)
	if err != nil {
		return err
	}

	result = validator.New()

	result.RegisterStructValidation(validatable.StructLevelValidation, validatable)

	err = result.Struct(validatable)
	if err != nil {
		return err
	}

	return nil
}
