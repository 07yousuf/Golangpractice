package main

import (
	"fmt"
)

type Human struct {
	Firstname string
	Lastname  string
	Age       int
	Country   string
}
type DomesticAnimal interface {
	RecieveAffection(from Human)
	GiveAffection(to Human)
}
type Cat struct {
	Name string
}
type Dog struct {
	Name string
}

func (c *Cat) RecieveAffection(from Human) {
	fmt.Printf("The cat named %s has received affection from Human named %s\n", c.Name, from.Firstname)
}
func (c *Cat) GiveAffection(to Human) {
	fmt.Printf("The cat named %s has received affection from Human named %s\n", c.Name, to.Firstname)
}
func (d *Dog) RecieveAffection(from Human) {
	fmt.Printf("The dog named %s has received affection from Human named %s\n", d.Name, from.Firstname)
}
func (d *Dog) GiveAffection(to Human) {
	fmt.Printf("The dog named %s has received affection from Human named %s\n", d.Name, to.Firstname)
}
func Pet(Animal DomesticAnimal, human Human) {
	Animal.RecieveAffection(human)
	Animal.GiveAffection(human)
}

func main() {
	// h := Human{
	// 	Firstname: "Abidul Islam",
	// 	Lastname:  "Yousuf",
	// 	Age:       21,
	// 	Country:   "Bangladesh",
	// }
	// c := Cat{
	// 	Name: "Meeu",
	// }
	// d := Dog{
	// 	Name: "BhogBhog",
	// }
	var h Human
	h.Firstname = "Yousuf"

	var c Cat
	c.Name = "Meeu"

	var d Dog
	d.Name = "Bhogbhog"

	Pet(&c, h)
	Pet(&d, h)
}
