package utils

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterValidation("mobile", validateMobile)
	validate.RegisterValidation("pan", validatePAN)
}

func ValidateStruct(data interface{}) map[string]string {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		fieldName := err.Field()
		//tag := err.Tag()
		errors[fieldName] = fmt.Sprintf("Validation failed on '%s' field", fieldName)
	}

	return errors
}

func validateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	re := regexp.MustCompile(`^[6-9]\d{9}$`)
	return re.MatchString(mobile)
}

func validatePAN(fl validator.FieldLevel) bool {
	pan := fl.Field().String()
	re := regexp.MustCompile(`^[A-Z]{5}[0-9]{4}[A-Z]$`) // regex: 5 letters, 4 digits, 1 letter
	return re.MatchString(pan)
}
