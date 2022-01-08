package maps

import "fmt"


func CreateMap(){
	m := make(map[string]int)

	m["tyes"] = 1

    fmt.Println(m)
}