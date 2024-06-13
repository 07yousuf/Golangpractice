package main
import "fmt"

type Shape interface{
  area() float64
}
type circle struct{
  radius float64 
}
type square struct{
  length float64 
}

func(c circle) area() float64 {
  return 3.14 * c.radius *c.radius 
}

func(s square) area() float64{
  area := s.length * s.length 
  return area 
}

func calculation(arealist []Shape){
  for _,elements := range arealist{
    fmt.Println("Area of Shape- ",elements.area())
  }
}
func main(){
  c := circle{radius: 5.00}
  s := square{length: 9.00}
  
  shape := []Shape{c,s}
  calculation(shape)
}
