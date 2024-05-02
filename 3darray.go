package main
import "fmt"

func main(){
    var cube [4][3][3] int
    for i:=0; i<4; i++{
        for j:=0; j<3; j++{
            for k:=0; k<3; k++{
                cube[i][j][k]=i + j + k
            }
        }
    }
    for i:=0; i<4; i++{
        for j:=0; j<3; j++{
            for k:=0; k<3; k++{
                fmt.Print(cube[i][j][k]," ")
            }
            fmt.Println()
        }
        fmt.Println()
  }
  
}

