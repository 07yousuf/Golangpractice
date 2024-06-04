package main
import (
  "fmt"
)
func main(){
  first := map[string]int{"a":1,"b":2,"c":3}
  fmt.Println("First -",first)
  second := map[string]int{"a": 1, "e": 5, "c": 3, "d": 4}
  fmt.Println("Second -",second)

  fmt.Println("After merged to first -")
  for key,value := range second{
    first[key] = value
  }
  fmt.Println(first)
}
