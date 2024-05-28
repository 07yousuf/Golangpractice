package main
import "fmt"

func main(){
  s := []string{"Bangladesh","India","Pakistan","Iran","Indonesia"}
  fmt.Println("before removing:\n",s)
  s2 := removeIndex(s, 2)
  fmt.Println("After remove:\n",s2)
}
func removeIndex(s []string, index int) []string{
  return append(s[:2],s[index+1:]...)
  }
