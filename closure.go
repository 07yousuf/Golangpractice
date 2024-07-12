package main 
import "fmt"

func greet(name string){

  var displayName = func() {
    fmt.Println("Hi ",name)
  }
  displayName()
}

func main(){
  var name string
  fmt.Println("One word name: ")
  fmt.Scan(&name)
  greet(name)

}
