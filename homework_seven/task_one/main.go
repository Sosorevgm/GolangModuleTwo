package main

import (
	"fmt"
	"reflect"
)

func main() {

	in := struct {
		Name    string
		Age     int
		Address string
		State   float64
	}{
		Name:    "Alex",
		Age:     25,
		Address: "Pokrovka",
		State:   1000.15,
	}

	values := map[string]interface{}{
		"Name":    "Joe",
		"Surname": "Biden",
		"Age":     78,
		"Dog":     "Jack",
		"State":   100000.58,
	}

	fmt.Println(in)
	err := Foo(&in, values)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(in)
}

func Foo(in interface{}, values map[string]interface{}) error {
	if in == nil {
		return fmt.Errorf("nil input interface")
	}
	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	} else {
		return fmt.Errorf("unaddressable input")
	}
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("input type is not the struct")
	}

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		if typeField.Type.Kind() == reflect.Struct {
			continue
		}

		for mKey, mValue := range values {
			if mKey == typeField.Name {
				if val.Field(i).Kind() == reflect.TypeOf(mValue).Kind() {
					val.Field(i).Set(reflect.ValueOf(mValue))
				}
			}
		}
	}
	return nil
}
