package main 
import "fmt"

func greet() func() int {

  number := 0
  return func() int {
    number++
    return number
  }

}

func main(){
  a := greet()
  for i:=0; i<3; i++ {
    fmt.Println(a())
  }

  a2 := greet()
  fmt.Println(a2())
  fmt.Println(a2())

}
