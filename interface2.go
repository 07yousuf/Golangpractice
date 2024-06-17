package main
import "fmt"

type employee interface{
  showInfo()
}

type identity struct{
  employee_id int 
}

type salary struct{
  salary_cal(basic int, tax int){
    var salary = (basic * tax) / 100
	return basic - salary
  }
}


func main(){
  obj1 := identity{employee_id: 12345}
  obj2 := salary{salary_cal(25000, 5)}

}
