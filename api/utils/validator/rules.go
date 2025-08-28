package validator

import "fmt"

type validator map[string]string

func New() validator {
	return validator{}
}

func (v validator) Required(value interface{}, field string) {
	if value == "" || value == nil {
		v[field] = fmt.Sprintf("The %s field is required.", field)
	}
}

func (v validator) Errors() bool {
	return len(v) > 0
}