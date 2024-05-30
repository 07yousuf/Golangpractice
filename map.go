package main
import "fmt"
func main(){
  employee := map[string]int {"mark":10,"scandy":20}
  fmt.Println(employee)
  
  var employeelist = make(map[string]int) //empty map 
  fmt.Println(employeelist)
  //determining how many item a map has by using len() function
  fmt.Println(len(employee))
  fmt.Println(len(employeelist))
}
