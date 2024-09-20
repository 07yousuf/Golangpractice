package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID   uint64
	Name string
	Sku  string
	Cat  Category
}
type Category struct {
	ID   uint64
	Name string
}

func main() {
	Stat := Product{ID: 34, Name: "Tea Pot", Sku: "Tp34", Cat: Category{ID: 3, Name: "Tea"}}
	a, err := json.Marshal(Stat)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(a))
}
