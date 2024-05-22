package main
import (
  "fmt"
  "strconv"
  "reflect"
)
func main(){
  sample := "250"
  newsample,err := strconv.Atoi(sample)
  if err != nil {
    fmt.Println("Converted falied\n")
  }else {
    fmt.Println("Converted to Integer\n")
  }
  fmt.Println(newsample,err,reflect.TypeOf(newsample))
}

