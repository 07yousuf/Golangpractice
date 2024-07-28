package main 
import (
  "fmt"
  set "github.com/deckarep/golang-set"
)

func main(){
  yellowfruit := set.NewSet()
  yellowfruit.Add("banana")
  yellowfruit.Add("lemon")
  yellowfruit.Add("pineapple")
  fmt.Println(yellowfruit)

  redfruit := set.NewSetFromSlice(
    []interface{}{"apple","cherry","strawberry"})
  fmt.Println(redfruit)

  fruit := yellowfruit.Union(redfruit)
  fmt.Println(fruit)
}
