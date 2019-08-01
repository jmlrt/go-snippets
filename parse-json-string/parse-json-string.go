// inspired from https://stackoverflow.com/questions/47270595/golang-parse-json-string-to-struct/47270657

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	JobTitle  string `json:"job"`
}

func main() {
	s := string(`{
		"firstName": "John",
		"lastName": "Smith",
		"age": 27,
		"job": "Engineer"
	  }`)
	data := person{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("First Name: %s\n", data.FirstName)
	fmt.Printf("Last Name: %s\n", data.LastName)
	fmt.Printf("Age: %d\n", data.Age)
	fmt.Printf("Job Title: %s\n", data.JobTitle)
}
