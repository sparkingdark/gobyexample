package threadgo

import "fmt"

func Mythread(a,b int) int{
	sum := make(chan int,1)
    go func(c,d int){
		fmt.Println("go thread")
		sum <- c+d
	}(a,b)
    fmt.Println("exited from go thread")
	return <-sum
    	
}