package main
import "fmt"

type rectangle func(int,int) int

type rectanglepara struct{
   length int
   breadth int 
   color string

  rect rectangle
}

func main(){
  var l int 
  var b int
  var c string 
  fmt.Print("Take value of LENGTH BREADTH & COLOR: \n")
  fmt.Scan(&l,&b,&c)
  sample := rectanglepara {
    length: l,
    breadth: b,
    color: c,

    rect: func(length int, breadth int) int{
      return length * breadth
    },
  }
  fmt.Println("Rectangle Area- ",sample.rect(sample.length,sample.breadth)) 
  fmt.Println("Color is- ",sample.color)
}
