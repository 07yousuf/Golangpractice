package main
import "fmt"
func main(){
  fmt.Println("Area:",rectangle(30,45))
}
func rectangle(l int, b int)(area int){
  parameter := 2 *(l+b)
  fmt.Println("Parameter:",parameter)
  area = l*b 
  return
  }
