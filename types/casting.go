package types

import "fmt"

func TypeCasting()  {

	var a interface{}
	a = 1
	fmt.Println(a.(string))
	a = true
	fmt.Println(a.(string))
	a = []int{1,2,3,4,5}
	fmt.Println(a.(string))
	a = struct {
		s string
		t int
	}{s: "sss", t:10}
	fmt.Println(a.(string))


}
