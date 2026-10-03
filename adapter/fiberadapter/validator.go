package fiberadapter

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type StructValidator struct {
	v *validator.Validate
}

func NewStructValidator() *StructValidator {
	v := validator.New(validator.WithRequiredStructEnabled())

	v.RegisterTagNameFunc(func(f reflect.StructField) string {
		for _, key := range []string{"json", "query"} {
			name := strings.SplitN(f.Tag.Get(key), ",", 2)[0]
			if name != "" && name != "-" {
				return name
			}
		}
		return f.Name
	})
	return &StructValidator{v: v}
}

func (s *StructValidator) Validate(out any) error {
	return s.v.Struct(out)
}

func validationFields(err error) (map[string]string, bool) {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return nil, false
	}

	fields := make(map[string]string, len(ve))
	for _, fe := range ve {
		fields[fe.Field()] = fieldMessage(fe)
	}
	return fields, true
}

func fieldMessage(fe validator.FieldError) string {
	field := fe.Field()
	unit := ""
	if fe.Kind() == reflect.String {
		unit = " characters"
	}

	switch fe.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email"
	case "numeric":
		return field + " must contain only digits"
	case "len":
		return field + " must be exactly " + fe.Param() + unit
	case "min":
		return field + " must be at least " + fe.Param() + unit
	case "max":
		return field + " must be at most " + fe.Param() + unit
	case "oneof":
		return field + " must be one of: " + strings.ReplaceAll(fe.Param(), " ", ", ")
	default:
		return field + " is invalid"
	}
}
