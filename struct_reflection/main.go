package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := MyAwesomeStruct{
		FieldString: "new line",
		FieldBool:   false,
		FieldFloat:  107.207,
	}
	thisMap := map[string]interface{}{
		"FieldString": "changed line",
		"FieldBool":   true,
		"FieldFloat":  0.0,
	}
	PrintStruct(v)
	MakeChanges(&v, thisMap)
	PrintStruct(v)
}

// MyAwesomeStruct is custom struct
type MyAwesomeStruct struct {
	FieldString string  `json:"field_string"`
	FieldBool   bool    `json:"field_bool"`
	FieldFloat  float64 `json:"field_float"`
}

// PrintStruct print custom struct
func PrintStruct(in interface{}) {
	if in == nil {
		return
	}

	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)

		fmt.Printf("\tname=%s, type=%s, value=%v, tag=`%s`\n",
			typeField.Name,
			typeField.Type,
			val.Field(i),
			typeField.Tag,
		)
	}
}

// MakeChanges does some magic with MyAwesomeStruct
func MakeChanges(in interface{}, values map[string]interface{}) error {
	if in == nil {
		return fmt.Errorf("Error occurred. Struct %s is nil. ", in)
	}

	if values == nil {
		return fmt.Errorf("Error occurred. Map %s is nil. ", values)
	}

	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("Error occurred. Input %s is not MyAwesomeStruct. ", val)
	}

	for key, value := range values {
		f := val.FieldByName(key)
		fmt.Println(f.Kind().String())
		switch f.Kind() {
		case reflect.Bool:
			v, ok := value.(bool)
			if !ok {
				return fmt.Errorf("Error occurred. Type assertion failed. %v", ok)
			}
			val.FieldByName(key).SetBool(v)
		case reflect.String:
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("Error occurred. Type assertion failed. %v", ok)
			}
			val.FieldByName(key).SetString(v)
		case reflect.Float64:
			v, ok := value.(float64)
			if !ok {
				return fmt.Errorf("Error occurred. Type assertion failed. %v", ok)
			}
			val.FieldByName(key).SetFloat(v)
		default:
			fmt.Println("Error occurred. Bad map.")
		}
	}
	return nil
}
