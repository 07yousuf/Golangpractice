package main
import "fmt"
func main(){
  for i:=10; i<=70; i++{
    if i!=70{
      fmt.Print(i,"+")
    } else {
      fmt.Println(i)
    }
    i+=5
  }
}
