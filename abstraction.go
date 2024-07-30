package main
import "fmt"

type shape interface{
  area() float32 
}

type rectangle struct{
  length float32
  width float32
}
type circle struct{
  radius float32
}
type triangle struct{
  length float32
  width float32
}
type square struct{
  length float32
}

func (r rectangle) area() float32 {
  return r.length*r.width
}
func (c circle) area() float32 {
  return 3.1416 * c.radius*c.radius
}
func (t triangle) area() float32 {
  return 0.5 * t.length * t.width
}
func (s square) area() float32 {
  return s.length * s.length
}

func show (s shape){
  result := s.area()
  fmt.Println(result)
}

func main (){
  r := rectangle{length:12,width:6}
  c := circle{radius:5}
  t := triangle{length:14,width:7}
  s := square{length:13}

  //var a shape
  show(r)
  show(c)
  show(t)
  show(s)
}
