package main
import "fmt"

func main(){
  array := [10]string {"a","b","c","d","e","f","g","h","i","j"}
  var aslice,bslice []string
  
  aslice = array[:]
  fmt.Println(aslice)
  aslice = array[:5]
  fmt.Println(aslice)
  aslice = array[2:7]
  fmt.Println(aslice)
  aslice = array[4:]
  fmt.Println(aslice)
  bslice = aslice[2:]//now bslice will count from the last update of aslice value
  fmt.Println(bslice)
  bslice = aslice[:]
  fmt.Println(aslice)
  bslice = aslice[4:]
  fmt.Println(bslice)
}
