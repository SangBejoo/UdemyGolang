package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name   string `required:"true" max:"10"`
	Age    int    `required:"true" max:"10"`
	Alamat string `required:"true" max:"10"`
}

func IsValid(data interface{}) (result bool) {
	t := reflect.TypeOf(data)
	result = true
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			isi := reflect.ValueOf(data).Field(i).Interface()
			result = isi != ""
			if result == false {
				return result
			}
		}
	}
	return true
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("type : ", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		structField := valueType.Field(i)
		fmt.Println("field : ", structField.Name, "With Type", valueType.Field(i).Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

func main() {
	readField(Sample{"Paul"})
	readField(Person{"Jack", 34, "Jack"})

	person := Person{"Jack", 33, ""}
	fmt.Println(IsValid(person))
}
