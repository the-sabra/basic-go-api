package util

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type (
	Validator struct {
		Instance  *validator.Validate
	}
	 
	ValidationErrorResponse struct {
		Errors []ValidationErrorDetail `json:"errors"`
		Status bool `json:"status"` 
	} 
	 
	ValidationErrorDetail struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}
)


func (cv *Validator) Validate(i interface{}) error {
	if err := cv.Instance.Struct(i); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errors []ValidationErrorDetail
			for _, err := range validationErrors {
				errors = append(errors, ValidationErrorDetail{
					Field:   err.Field(),
					Message: tagMessages(err),
				})
			}
			return echo.NewHTTPError(http.StatusBadRequest, ValidationErrorResponse{Errors: errors,Status: false})
		} 
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}      

func tagMessages(fe validator.FieldError) string {
	var message string
	switch fe.Tag() {
	case "required":
		message = "This field is required"
	case "email":
		message = "Invalid email format"
	case "min":
		message = "Value is too short"
	case "max":
		message = "Value is too long"
	case "len":
		message = "Invalid length"
	default:
		message = "Invalid value"
	}
	return message
}  