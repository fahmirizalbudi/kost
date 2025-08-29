package validator

import (
	"fmt"
	"reflect"
)

type validator map[string]string

func New() validator {
	return validator{}
}

func (v validator) Required(value any, field string) {
	if value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil()) || fmt.Sprint(value) == "" {
		v[field] = fmt.Sprintf("The %s field is required.", field)
	}
}

func (v validator) Errors() bool {
	return len(v) > 0
}