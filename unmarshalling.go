package main

import (
	"encoding/json"
	"fmt"
)

type Json struct {
	cat `json:"cat"`
}
type cat struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonstore := []byte(`{"cat":{"name":"stroki","age":3}}`)
	variable := Json{}
	err := json.Unmarshal(jsonstore, &variable)
	if err != nil {
		panic(err)
	}
	fmt.Println(variable.cat.Name)
	fmt.Println(variable.cat.Age)
}
