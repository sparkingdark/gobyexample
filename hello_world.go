package main 

import (
	"fmt"
	values "github.com/gobyexample/values"
    variables "github.com/gobyexample/variables"
	constant "github.com/gobyexample/constant"
	forloop "github.com/gobyexample/forloop"
    ifelse "github.com/gobyexample/ifelse"
	Switch "github.com/gobyexample/Switch"
)

func main(){
	fmt.Println("Hello world")
    values.Values()
	variables.TakeVariables(1,2)
	constant.Declareconstant()
    forloop.Forloop()
	ifelse.Ifelse(1)
	Switch.Switch(2)
}