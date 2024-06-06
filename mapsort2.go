package main
import (
  "fmt"
  "sort"
)
func main(){
  unsorted := map[string]int{"Bangladesh":35,"Finland":90,"Slovakia":55}
  fmt.Println("Before Sorting countries value-",unsorted)

  valueslice := make([]int,0,len(unsorted))

  for _,value := range unsorted{
    valueslice = append(valueslice,value)
  }

  sort.Ints(valueslice)

  fmt.Println("After sorting-")
  for _,v := range valueslice{
    fmt.Println(v)
  }
}
