package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

var jsonData = `
{
	"name": "John Doe",
	"age": 30,
	"isProgrammer": true,
	"skills": ["Go", "Python", "JavaScript"],
	"address": {
		"street": "123 Main St",
		"city": "Anytown"
	}
}`

/*
1. unmarshal for []byte
2. decode for io.reader
*/
func Decode() {
	result := map[string]any{}
	// err := json.Unmarshal([]byte(jsonData), &result)
	// if err != nil {
	// 	log.Fatalf("Error during Unmarshal: %v", err)
	// }

	reader := bytes.NewReader([]byte(jsonData))
	err := json.NewDecoder(reader).Decode(&result)
	if err != nil {
		log.Fatalf("Error during NewDecoder: %v", err)
	}
	fmt.Printf("%+v\n", result)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`

	// omitempty skips the field if it's empty when encoding.
	Email  string `json:"email,omitempty"`
	Email2 string `json:"email2"`
}

func Encode() {
	person := Person{
		Name: "John Doe",
		Age:  30,
	}
	raw, err := json.Marshal(person)
	if err != nil {
		log.Fatalf("Error during Marshal: %v", err)
	}
	fmt.Println(string(raw))
}
