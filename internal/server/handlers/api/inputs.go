package api

import "reflect"

func makeInputsMap(data any) map[string]string {
	dataType := reflect.TypeOf(data)
	if dataType.Kind() != reflect.Struct {
		return nil
	}

	dataValues := reflect.ValueOf(data)
	inputs := make(map[string]string, dataType.NumField())
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		value := dataValues.Field(i)
		if name := field.Tag.Get("form"); name != "" {
			inputs[field.Name] = value.String()
		}
	}

	return inputs
}
