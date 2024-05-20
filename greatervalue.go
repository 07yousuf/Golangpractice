package main

import "fmt"

func main(){
  var a,b,c int
  fmt.Println("Give your value of a,b & c:\n")
  fmt.Scan(&a,&b,&c) 

  Detector(a,b,c)
}

func Detector(a1,b2,c3 int){
  if a1>b2 && a1>c3 {
    fmt.Println("Greater Value is:",a1)
  } else if b2>a1 && b2>c3{
    fmt.Println("Greater Value is:",b2)
  } else if c3>a1 && c3>b2{
    fmt.Println("Greater Value is:",c3)
  } else {
   fmt.Println("There is no single greatest value\n")
  }
}

