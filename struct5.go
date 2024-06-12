package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
	color   string
}
func main(){
  r1 := rectangle{40,20,"Green"}
  r2 := r1
  r2.color = "black"
  fmt.Println(r2)
  fmt.Println(r1)

  r3 := &r1
  r3.color = "red"
  fmt.Println(r3)
  fmt.Println(r1)

}
