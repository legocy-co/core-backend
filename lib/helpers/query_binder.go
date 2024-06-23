package helpers

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// BindQueryParamsToStruct dynamically populates the fields of a struct based on url.Values.
func BindQueryParamsToStruct(v interface{}, values url.Values) error {
	if reflect.TypeOf(v).Kind() != reflect.Ptr || reflect.ValueOf(v).Elem().Kind() != reflect.Struct {
		return errors.New("target must be a pointer to a struct")
	}

	val := reflect.ValueOf(v).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := field.Type()
		fieldTag := typ.Field(i).Tag
		formTag := fieldTag.Get("form")

		if formTag == "" {
			continue
		}

		switch fieldType.Kind() {
		case reflect.Ptr:
			ptrFieldType := fieldType.Elem()
			if ptrFieldType.Kind() == reflect.Struct {
				// Handle pointer to nested struct
				newFieldPtr := reflect.New(ptrFieldType)
				if err := BindQueryParamsToStruct(newFieldPtr.Interface(), extractSubValues(values, formTag)); err != nil {
					return err
				}
				if !newFieldPtr.Elem().IsZero() {
					field.Set(newFieldPtr)
				}
			} else {
				// Handle pointer to primitive type
				handlePointerField(field, values.Get(formTag))
			}
		case reflect.Struct:
			// Handle nested struct
			newFieldStruct := reflect.New(fieldType).Elem()
			if err := BindQueryParamsToStruct(newFieldStruct.Addr().Interface(), extractSubValues(values, formTag)); err != nil {
				return err
			}
			field.Set(newFieldStruct)
		case reflect.Slice:
			handleSliceField(field, values[formTag])
		default:
			setFieldValue(field, values.Get(formTag))
		}
	}

	return nil
}

// Existing helper functions (handlePointerField, handleSliceField, setFieldValue) should be included here.

// extractSubValues extracts sub-struct values from url.Values for nested structures.
func extractSubValues(values url.Values, prefix string) url.Values {
	subValues := url.Values{}
	for key, value := range values {
		if strings.HasPrefix(key, prefix+"[") && strings.HasSuffix(key, "]") {
			subKey := key[len(prefix)+1 : len(key)-1] // Remove the prefix and the enclosing brackets
			subValues[subKey] = value
		}
	}

	return subValues
}

// handlePointerField sets the field for pointer types
func handlePointerField(field reflect.Value, value string) {
	if value == "" {
		return
	}

	newField := reflect.New(field.Type().Elem())

	setFieldValue(newField.Elem(), value)
	field.Set(newField)
}

func handleSliceField(field reflect.Value, values []string) {
	elementType := field.Type().Elem()

	slice := reflect.MakeSlice(field.Type(), 0, len(values))

	for _, value := range values {
		newElement := reflect.New(elementType).Elem()
		if trySetSliceElement(newElement, value) {
			slice = reflect.Append(slice, newElement)
		}
	}

	if slice.Len() == 0 {
		slice = reflect.Zero(field.Type())
	}

	field.Set(slice)
}

func trySetSliceElement(element reflect.Value, value string) bool {
	switch element.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if intValue, err := strconv.ParseInt(value, 10, element.Type().Bits()); err == nil {
			element.SetInt(intValue)
			return true
		} else {
			return false
		}
	case reflect.Float32, reflect.Float64:
		if floatValue, err := strconv.ParseFloat(value, element.Type().Bits()); err == nil {
			element.SetFloat(floatValue)
			return true
		} else {
			return false
		}
	case reflect.Bool:
		if boolValue, err := strconv.ParseBool(value); err == nil {
			element.SetBool(boolValue)
			return true
		} else {
			return false
		}
	case reflect.String:
		element.SetString(value)
		return true
	}

	return false
}

// Modify setFieldValue to return false if parsing fails.
func setFieldValue(field reflect.Value, value string) bool {
	if value == "" {
		return false
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			field.SetInt(intValue)
		} else {
			return false
		}
	case reflect.Float32, reflect.Float64:
		if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			field.SetFloat(floatValue)
		} else {
			return false
		}
	case reflect.Bool:
		if boolValue, err := strconv.ParseBool(value); err == nil {
			field.SetBool(boolValue)
		} else {
			return false
		}
	default:
		return false
	}
	return true
}
