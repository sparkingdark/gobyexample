package constant

import (
	"fmt"
	"math"
)

func Declareconstant() {
	const n = 1200000

	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
	
}
