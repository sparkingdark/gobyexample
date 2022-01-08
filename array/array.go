package array

import "fmt"

func MyArray(n int){
	a:= make([]int,n)

	for i:=0;i<4;i++{
		a[i] = i
	}

	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}