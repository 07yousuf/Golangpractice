package main
import (
        "fmt"
        "reflect"
      )

func main(){
  s1 := []string {"Bangladesh","India","Pakistan","iran","Turkey"}

  fmt.Println(itemexists(s1,"Turkey"))
  fmt.Println(itemexists(s1,"Africa"))
}

func itemexists(s2 interface{}, item interface{}) bool{
  s := reflect.ValueOf(s2)

  if s.Kind() != reflect.Slice {
    panic("Invalid data type")
  }

  for i:=0; i<s.Len(); i++ {
    if s.Index(i).Interface() == item {
      return true
    }
  }
  return false
}
