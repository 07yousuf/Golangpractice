package main
import "fmt"

func main(){
  var family = make(map[string]int)
  family["yousuf"] = 58
  family["rakiun"] = 25
  family["asif"] = 60
  family["sayma"] = 50

  fmt.Println(family["rakiun"])
  //removing map item
  delete(family,"sayma")
  fmt.Println(family)
}
