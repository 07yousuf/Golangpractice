package main
import "fmt"

type employee struct{
  name string
  age int
}

func (obj *employee) Infoinitial(){
  if obj.name =="" {
    obj.name="random name"
  }
  if obj.age ==0 {
    obj.age=20
  }
}
func main(){
  emp := employee{name:"Abidul Islam Yousuf"}
  emp.Infoinitial()
  fmt.Println(emp)

  emp2 := employee{age: 30}
  emp2.Infoinitial()
  fmt.Println(emp2)
}
