package helpers

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

type NestedQueryBinder struct{}

func (b *NestedQueryBinder) BindQuery(c *gin.Context, obj any) error {
	// Get the query parameters directly from gin.Context.
	values := c.Request.URL.Query()

	// Reflect on the object passed to the function to be able to set its fields.
	objVal := reflect.ValueOf(obj).Elem()

	// Iterate over each field in the struct.
	for i := 0; i < objVal.NumField(); i++ {
		field := objVal.Field(i)
		fieldType := objVal.Type().Field(i)

		// Get the form tag for the current field. If not set, skip this field.
		formTag := fieldType.Tag.Get("form")
		if formTag == "" {
			continue
		}

		// Check if the current field is a struct, indicating a nested structure.
		if fieldType.Type.Kind() == reflect.Struct {
			// For nested structs, iterate over each nested field.
			for j := 0; j < field.NumField(); j++ {
				nestedField := field.Field(j)
				nestedFieldType := fieldType.Type.Field(j)

				// Get the form tag for the nested field.
				nestedFormTag := nestedFieldType.Tag.Get("form")

				// Construct the query parameter key using the prefix format (e.g., "b__field1").
				queryParam := formTag + "__" + nestedFormTag

				// If the query parameter exists, bind its value to the nested field.
				if value, exists := values[queryParam]; exists && len(value) > 0 {
					setValue(nestedField, nestedFieldType.Type.Kind(), value[0])
				}
			}
		} else {
			// For non-struct fields, directly bind the query parameter to the field if it exists.
			if value, exists := values[formTag]; exists && len(value) > 0 {
				setValue(field, fieldType.Type.Kind(), value[0])
			}
		}
	}
	return nil
}

// setValue sets the value of a field based on its kind.
func setValue(field reflect.Value, kind reflect.Kind, value string) {
	switch kind {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			field.SetInt(intValue)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if uintValue, err := strconv.ParseUint(value, 10, 64); err == nil {
			field.SetUint(uintValue)
		}
	case reflect.Bool:
		if boolValue, err := strconv.ParseBool(value); err == nil {
			field.SetBool(boolValue)
		}
	// Add more types as needed
	default:
		// Unsupported type
	}
}
