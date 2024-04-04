package helpers

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// BindQueryParamsToStruct dynamically binds query params to struct fields.
func BindQueryParamsToStruct(dst interface{}, values url.Values) {
	v := reflect.ValueOf(dst).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		typeField := v.Type().Field(i)
		queryKey := typeField.Tag.Get("form")

		if queryKey == "" {
			continue
		}

		if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Struct {
			// Prepare sub-values for the nested struct based on its prefix.
			subValues := filterValuesForNestedStruct(queryKey, values)
			if len(subValues) > 0 {
				// Only instantiate and populate the nested struct if relevant parameters exist.
				nestedStructPtr := reflect.New(field.Type().Elem())
				BindQueryParamsToStruct(nestedStructPtr.Interface(), subValues)
				field.Set(nestedStructPtr)
			}
		} else if field.Kind() == reflect.Slice {
			sliceHandler(field, typeField, values)
		} else {
			// Handle direct assignment for basic types and pointers.
			value, exists := values[queryKey]
			if exists && len(value) > 0 {
				assignValueToField(field, value[0])
			}
		}
	}
}

func filterValuesForNestedStruct(prefix string, values url.Values) url.Values {
	subValues := make(url.Values)
	for k, v := range values {
		if strings.HasPrefix(k, prefix+"[") && strings.HasSuffix(k, "]") {
			subKey := strings.TrimPrefix(k, prefix+"[")
			subKey = strings.TrimSuffix(subKey, "]")
			subValues[subKey] = v
		}
	}
	return subValues
}

func sliceHandler(field reflect.Value, typeField reflect.StructField, values url.Values) {
	queryKey := typeField.Tag.Get("form")
	if sliceValues, exists := values[queryKey]; exists {
		sliceType := field.Type().Elem()
		slice := reflect.MakeSlice(reflect.SliceOf(sliceType), 0, len(sliceValues))
		for _, value := range sliceValues {
			tempValue := reflect.New(sliceType).Elem()
			if assignValueToField(tempValue, value) {
				slice = reflect.Append(slice, tempValue)
			}
		}
		field.Set(slice)
	}
}

func assignValueToField(field reflect.Value, value string) bool {
	if field.Kind() == reflect.Ptr {
		// Allocate memory for pointers
		field.Set(reflect.New(field.Type().Elem()))
		field = field.Elem()
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int64:
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			field.SetInt(intValue)
			return true
		}
	case reflect.Float32, reflect.Float64:
		if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			field.SetFloat(floatValue)
			return true
		}
	case reflect.Bool:
		if boolValue, err := strconv.ParseBool(value); err == nil {
			field.SetBool(boolValue)
			return true
		}
	}
	return false // Unsupported type
}
