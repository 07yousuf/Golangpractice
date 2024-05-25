package main
import "fmt"

func main(){
  a := make([]int,2,5)
  a[0]=10
  a[1]=20
  fmt.Println(a)
  fmt.Printf("Length is %d capacity is %d\n",len(a),cap(a))

  a = append(a,50,20,45,33,47)
  fmt.Println(a)
  fmt.Printf("Now Length is %d capacity is %d\n",len(a),cap(a))
}
