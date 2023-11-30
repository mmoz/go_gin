package validate

import (
	"errors"
	"log"

	"github.com/go-playground/validator"
)

func ValidateStruct(req any) error {
	v := validator.New()
	err := v.Struct(req)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println("Invalid validation error")

			return errors.New("Invalid validation error")
		}
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			tag := err.Tag()
			if tag == "required" {
				log.Printf("%s is required\n", fieldName)
				return errors.New(fieldName + " is required")
			}
		}
	}
	return nil
}
