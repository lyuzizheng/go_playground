package pointers

import (
	"fmt"
)


//go:noinline
func UpdateMap(m map[string]int) {
	m["one"] = 0
}


func NewMap(){
	var m = map[string]int{"one": 1, "two":2}

	fmt.Printf("NewMap: %v\n", m)
	UpdateMap(m)
	fmt.Printf("AfterReturn: %v\n", m)
}