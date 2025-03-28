package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// BuildValidationErrors maps validator tags to user-friendly English messages
func BuildValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	if verrs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range verrs {
			field := e.Field()
			switch e.Tag() {
			case "required":
				errors[field] = "This field is required"
			case "min":
				errors[field] = fmt.Sprintf("Minimum length is %s characters", e.Param())
			case "max":
				errors[field] = fmt.Sprintf("Maximum length is %s characters", e.Param())
			case "gte":
				errors[field] = fmt.Sprintf("Must be greater than or equal to %s", e.Param())
			case "lte":
				errors[field] = fmt.Sprintf("Must be less than or equal to %s", e.Param())
			case "email":
				errors[field] = "Invalid email format"
			case "uuid4":
				errors[field] = "Invalid UUID v4 format"
			case "url":
				errors[field] = "Invalid URL"
			case "alphanum":
				errors[field] = "Only letters and numbers are allowed"
			case "numeric":
				errors[field] = "Must be a numeric value"
			case "eq":
				errors[field] = fmt.Sprintf("Must be equal to %s", e.Param())
			case "ne":
				errors[field] = fmt.Sprintf("Must not be equal to %s", e.Param())
			case "len":
				errors[field] = fmt.Sprintf("Must be exactly %s characters long", e.Param())
			case "status":
				errors[field] = "Invalid status value"
			default:
				errors[field] = "Invalid value"
			}
		}
	}

	return errors
}
