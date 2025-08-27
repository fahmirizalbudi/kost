package validate

import "fmt"

func Required(validations map[string]string, value interface{}, field string) {
	if value == "" || value == nil {
		validations[field] = fmt.Sprintf("The %s field is required.", field)
	}
}