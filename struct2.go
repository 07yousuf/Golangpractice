package main
import "fmt"

type salary struct{
  basic,HRA,TA float64
}

type employee struct{
    firstname,lastname,email string
    monthlysalary []salary
  }
func main(){
  e := employee{
    firstname: "Abdul",
    lastname: "karim",
    email: "abdulkarim@gmail.com",
    
    monthlysalary: []salary{
      salary{
        basic: 20000.00,
        HRA: 6000.00,
        TA: 3000.00,
      },
      salary{
        basic: 25000.00,
        HRA: 11000.00,
        TA: 5000.00,
      },
    },
  }

  fmt.Println(e.firstname,e.lastname)
  fmt.Println(e.email)
  fmt.Println("Salary range of a employee-")
  fmt.Println(e.monthlysalary[0])
  fmt.Println(e.monthlysalary[1])
}
