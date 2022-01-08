package ranges

import "fmt"

func Ranges(){
	nums := []int{2, 3, 4}
    sum := 0
    for idx, num := range nums {
		fmt.Println("the index is:",idx)
        sum += num
    }
    fmt.Println("sum:", sum)
}