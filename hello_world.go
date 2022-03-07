package main

import (
	"fmt"

	Switch "github.com/gobyexample/Switch"
	anon "github.com/gobyexample/anonymous"
	array "github.com/gobyexample/array"
	slices "github.com/gobyexample/array/slices"
	channel "github.com/gobyexample/channel"
	constant "github.com/gobyexample/constant"
	errors "github.com/gobyexample/errors"
	forloop "github.com/gobyexample/forloop"
	ifelse "github.com/gobyexample/ifelse"
	maps "github.com/gobyexample/maps"
	multi "github.com/gobyexample/multiplereturnvalue"
	ranges "github.com/gobyexample/ranges"
	thread "github.com/gobyexample/threadgo"
	values "github.com/gobyexample/values"
	variables "github.com/gobyexample/variables"
	variad "github.com/gobyexample/variadic"
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
    
	fmt.Println(thread.Mythread(1,2))

}

