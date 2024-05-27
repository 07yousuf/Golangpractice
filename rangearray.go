package main
import "fmt"

func main(){

  array := [...]int{10,20,30,40,50}

  fmt.Println("Exmaple -1")
  for index,element := range array{
    fmt.Println(index,"=>",element)
  }

  fmt.Println("Example -2")
  for _,value := range array{
    fmt.Println(value)
  }

  fmt.Println("Example -3")
  j := 0
  for range array{
    fmt.Println(array[j])
    j++
  }
}
