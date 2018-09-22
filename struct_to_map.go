package mirror

import "reflect"

// StructToMap takes any struct instance and returns a map of strings (struct
// key names) to the values set in the struct.
func StructToMap(object interface{}) map[string]interface{} {
	objType := reflect.TypeOf(object)
	objValue := reflect.ValueOf(object)
	if objType.Kind() != reflect.Struct {
		return nil
	}

	result := make(map[string]interface{}, objType.NumField())
	for i := 0; i < objType.NumField(); i++ {
		fieldType := objType.Field(i)
		field := objValue.Field(i)
		if field.CanInterface() {
			result[fieldType.Name] = field.Interface()
		}
	}
	return result
}
