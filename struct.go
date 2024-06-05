package main
import "fmt"

type rectangle struct{
  length int
  breadth int
  color string
}
func main(){
  //way struct instantiation  
  A := new(rectangle)
  A.length = 20
  A.breadth = 30
  A.color = "green"
  fmt.Println(*A)
  //another way of struct instantiation
  B := &rectangle{}
  B.length = 25
  B.color = "red"
  fmt.Println(*B)
}
