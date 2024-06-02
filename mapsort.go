package main
import (
  "fmt"
  "sort"
)
func main(){
  unsorted := map[string]int{"Finland":90,"Bangladesh":35,"Slovakia":55}
  fmt.Println("Before keys ascending alphabetical order-",unsorted)

  slicekey := make([]string,0,len(unsorted))
  
  for k := range unsorted {
    slicekey = append(slicekey,k)
  }
  sort.Strings(slicekey)

  fmt.Println("After ascending alphabetical order-")
  for _,k := range slicekey{
    fmt.Println(k,unsorted[k])
  }
}

