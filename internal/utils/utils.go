package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func InspectStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("Not a struct")
	}

	fmt.Printf("Struct has %d fields:\n", t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fmt.Printf("  %d. Name: %s, Type: %v, Value: %v, Tag: %v\n",
			i+1, field.Name, field.Type, value.Interface(), field.Tag)
	}
	return nil
}

func GetStructTag(s interface{}) []string {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("Not a struct"))
	}

	fmt.Printf("Struct has %d fields:\n", t.NumField())
	var tags []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := string(field.Tag)
		tag = strings.Split(tag, ":")[1]
		tag = strings.Trim(tag, `"`)
		tags = append(tags, tag)
		fmt.Printf("  %d. Name: %s, Type: %v, Value: %v, Tag: %v\n",
			i+1, field.Name, field.Type, value.Interface(), field.Tag)
	}
	return tags
}
