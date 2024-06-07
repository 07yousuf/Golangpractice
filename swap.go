package main
import "fmt"

func swap(xptr *int, yptr *int){
  *xptr = 2
  *yptr = 1
}
func main(){
  x := 1
  y := 2
  fmt.Println("Before swaping x=",x,"y=",y)
  swap(&x,&y)
  fmt.Println("After swaping x=",x,"y=",y)
}
