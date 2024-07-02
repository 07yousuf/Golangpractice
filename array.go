package main
import "fmt"

func main(){
    var cube [5][4][3][3] int
    for h:=0; h<5; h++{
      for i:=0; i<4; i++{
        for j:=0; j<3; j++{
            for k:=0; k<3; k++{
                cube[h][i][j][k]=h+i + j + k
            }
        }
      }
    }
    for h:=0; h<5; h++{
      for i:=0; i<4; i++{
        for j:=0; j<3; j++{
            for k:=0; k<3; k++{
                fmt.Print(cube[h][i][j][k]," ")
            }
            fmt.Println()
        }
        fmt.Println()
      }
      fmt.Println()
    }
}

