package validate

import "fmt"

type Validator map[string]string

func New() Validator {
	return Validator{}
}

func (v Validator) Required(value interface{}, field string) {
	if value == "" || value == nil {
		v[field] = fmt.Sprintf("The %s field is required.", field)
	}
}

func (v Validator) Errors() bool {
	return len(v) > 0
}