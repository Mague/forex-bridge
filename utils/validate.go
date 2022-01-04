package utils

import validator "gopkg.in/go-playground/validator.v8"

var validate *validator.Validate

func init() {
	config := &validator.Config{TagName: "validate"}
	validate = validator.New(config)
}
func Email(str string) bool {

	errs := validate.Field(str, "required,email")
	if errs != nil {
		return false
	}
	return true
}
