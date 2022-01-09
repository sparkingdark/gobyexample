package channel

import "fmt"

func TestChannel(a chan bool){
	fmt.Println("print it")
	a <- true
}