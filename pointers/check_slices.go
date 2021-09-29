package pointers

import (
	"code.byted.org/im_cloud/playground/utils"
	"fmt"
	"reflect"
)


//go:noinline
func UpdateSlice(s []int) {
	utils.DumpObject("DeforeUpdateSlice", reflect.ValueOf(&s))
	fmt.Printf("DeforeUpdateSlice: %v\n", s)
	s[1] = 5
	s = append(s, 8)
	utils.DumpObject("UpdateSlice", reflect.ValueOf(&s))
	fmt.Printf("AfterUpdate: %v\n", s)
}


func CheckSlice(){
	var s []int
	utils.DumpObject("NewSlice", reflect.ValueOf(&s))
	fmt.Printf("NewSlice: %v\n", s)
	UpdateSlice(s)
	utils.DumpObject("AfterReturnSlice", reflect.ValueOf(&s))
	fmt.Printf("AfterReturn: %v\n", s)
}