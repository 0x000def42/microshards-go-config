package utils

import (
	"fmt"

	"github.com/0x000def42/microshards-go-config/models"
	"github.com/go-playground/validator"
)

// ValidationError wraps the validators FieldError so we do not
// expose this to out code
type ValidationError struct {
	validator.FieldError
}

type ValidationErrorMessages struct {
	Messages []string `json:"messages"`
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// Validation contains
type Validation struct {
	validate *validator.Validate
}

// NewValidation creates a new Validation type
func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("user_role", validateUserRole)
	return &Validation{validate}
}

func Validate(i interface{}) ValidationErrors {
	v := NewValidation()
	return v.Validate(i)
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	err := v.validate.Struct(i)
	if err == nil {
		return nil
	}

	errs := err.(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	var returnErrs []ValidationError
	for _, err := range errs {
		// cast the FieldError into our ValidationError and append to the slice
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}

// Custom validations

func validateUserRole(fl validator.FieldLevel) bool {
	field := fl.Field().String()

	if field == string(models.USER_ROLE_ADMIN) {
		return true
	}

	if field == string(models.USER_ROLE_USER) {
		return true
	}

	if field == string(models.USER_ROLE_GUEST) {
		return true
	}

	return false

}
