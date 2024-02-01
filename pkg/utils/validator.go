package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func GetJSONTag(obj interface{}, fieldName string) string {
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)

		if field.Name == fieldName {
			jsonTag := field.Tag.Get("json")

			// If the tag is empty, use the field name
			if jsonTag == "" {
				jsonTag = field.Name
			}

			return jsonTag
		}
	}

	return ""
}
