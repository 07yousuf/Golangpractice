package main 
import "fmt"

func determine(a int ,b int) int{
  fb := a 
  eb := b 
  collection := 0
  result1 := fb/eb 
  var i int 
  for i=result1; i>=eb; i--{

    collection += i
    i /= eb
    i++

  }
  collection += i
  total := fb + collection 
  return total 
}

func main(){
  var full_bottles int
  var exchange_bottles int
  fmt.Println("take input full_bottles and exchange_bottles : ")
  fmt.Scan(&full_bottles,&exchange_bottles)

  FinalResult := determine(full_bottles,exchange_bottles)

  fmt.Println("The number of total drinkable water bottles is: ",FinalResult)

}
