package  main
import "fmt"

func main(){
  var age int
  var name string
  fmt.Println("Input your Name & Age: ")
  fmt.Scan(&name,&age)
  fmt.Println("Before: ",name,age)

  update(&age,&name)
  
  fmt.Println("After: ",name,age)
}
func update(a *int, b *string){
  *a+=5
  *b+=" Yousuf"
  return//no need to return value, because pointer update my address value//in golang it's called pass by value not reference 
}
