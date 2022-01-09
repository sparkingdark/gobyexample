package main 

import (
	"fmt"
	values "github.com/gobyexample/values"
    variables "github.com/gobyexample/variables"
	constant "github.com/gobyexample/constant"
	forloop "github.com/gobyexample/forloop"
    ifelse "github.com/gobyexample/ifelse"
	Switch "github.com/gobyexample/Switch"
    array "github.com/gobyexample/array"
	slices "github.com/gobyexample/array/slices"
	maps "github.com/gobyexample/maps"
    ranges "github.com/gobyexample/ranges"
	multi "github.com/gobyexample/multiplereturnvalue"
	variad "github.com/gobyexample/variadic"
	errors "github.com/gobyexample/errors"
	anon "github.com/gobyexample/anonymous"
	channel "github.com/gobyexample/channel"
)

func main(){
	fmt.Println("Hello world")
    values.Values()
	variables.TakeVariables(1,2)
	constant.Declareconstant()
    forloop.Forloop()
	ifelse.Ifelse(1)
	Switch.Switch(2)
	array.MyArray(5)
	slices.MakeArray()
	maps.CreateMap()
	ranges.Ranges()
	fmt.Println(multi.Vals())
	variad.Sum([]int{1,2,3,4,5}...)
	fmt.Println(anon.Intseq()())
	fmt.Println(errors.F1(42))
	done := make(chan bool, 1)
	go channel.TestChannel(done)

	<-done


}

