package main

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {

	val := getValue(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {

		case reflect.String:
			fn(field.String())

		case reflect.Struct:
			walk(field.Interface(), fn)

		}
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

func main() {
	var res []string
	person := struct {
		Name string
		City string
	}{"Jo", "SPB"}

	walk(person, func(input string) {
		res = append(res, input)
	})

	fmt.Println(res)
}
