package main 
import "fmt"

type shape interface {
  determine() int
}
type numbers struct {
   num1 int 
   num2 int
}
func (n numbers) determine() int{
  if n.num1 > n.num2{
    return n.num1
  } else if n.num1 == n.num2{
    return 0
  } else{
    return n.num2
  }
}
func show (s shape){
  result := s.determine()
  fmt.Println(result)
}

func main(){
  var num1,num2 int
  fmt.Scan(&num1,&num2)
  instance := numbers{num1,num2}
  show(instance)
}
