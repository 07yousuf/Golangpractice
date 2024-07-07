package main
import "fmt"

func factorial(a int) int{
  var answer int  
  if a==0 {
    return 1
  } else {
      answer = a*factorial(a-1)
    }
    return answer
}

func main(){
  var number int
  fmt.Scan(&number)
  fmt.Println(factorial(number))
}
