package reflection

import (
	"fmt"
	"reflect"
)

// basic
// reflect.TypeOf:  get the type of a variable.
// reflect.ValueOf: get the value of a variable.

// for a struct
// reflect.Value.Field: 		get the value of a struct field.
// reflect.Value.FieldByName: 	get the value of a struct field by its name.
// reflect.Value.FieldByIndex: 	get the value of a struct field by its index.
// reflect.Value.Set: 			set the value of a variable.

// ref: https://medium.com/@alandev9751210/reflect-go-%E8%AA%9E%E8%A8%80%E7%9A%84%E9%8F%A1%E5%AD%90-efaaf16f329d

type Person struct {
	Name string
	Age  int
}

func Reflect() {
	person := Person{Name: "Alan", Age: 30} // Use a pointer to the struct
	val := reflect.ValueOf(&person).Elem()  // Elem(): return the pointer points to

	ageField := val.FieldByName("Age")
	fmt.Printf("Age: %v\n", ageField) // Age: 30

	val.FieldByName("Age").SetInt(31)
	fmt.Printf("Updated Age: %v\n", person.Age) // Updated Age: 31
}

// iterate all fields
func IterateByReflect() {
	p := Person{"Alan", 30}
	valueOfP := reflect.ValueOf(p)

	for i := range valueOfP.NumField() {
		field := valueOfP.Field(i)
		fmt.Printf("Field Name: %s, Field Value: %v\n", valueOfP.Type().Field(i).Name, field.Interface())
	}
}
