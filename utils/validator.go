package utils

import (
	"github.com/go-playground/validator/v10"
)

type RequestValidator struct {
	validator *validator.Validate
}

// RegisterValidator
// Register new echo validator
func RegisterValidator() *RequestValidator {
	return &RequestValidator{
		validator: validator.New(),
	}
}

// Validate
// Validate Request payloads
func (cv *RequestValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
