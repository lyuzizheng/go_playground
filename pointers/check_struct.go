//go:build exclude
// +build exclude

package pointers

import (
	"fmt"
)

type MyStruct struct {
	A int
	B string
}

//go:noinline
func UpdateStruct(s MyStruct) {
	s.A = 5
}

func NewStruct() {
	var s = MyStruct{
		A: 0,
		B: "0",
	}

	fmt.Printf("NewStruct: %v\n", s)
	UpdateStruct(s)
	fmt.Printf("UpdatedStruct: %v\n", s)
}

//func NewStruct(){
//	var s = MyStruct{
//		A: 0,
//		B: "0",
//	}
//	var t = s
//	utils.DumpObject("NewStruct1", reflect.ValueOf(&s))
//	utils.DumpObject("NewStruct2", reflect.ValueOf(&t))
//	fmt.Printf("NewStruct: %v\n", s)
//	UpdateStruct(&s)
//	utils.DumpObject("AfterReturnStruct1", reflect.ValueOf(&s))
//	utils.DumpObject("AfterReturnStruct2", reflect.ValueOf(&t))
//	fmt.Printf("AfterReturn: %v\n", s)
//}
