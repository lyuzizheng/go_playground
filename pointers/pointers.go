package pointers

import (
	"code.byted.org/im_cloud/playground/utils"
	"reflect"
)

func PointerDemo() {
	utils.PrintObjectDumpTableHeader()


	var b bool
	var i int
	var s string
	utils.DumpObject("b", reflect.ValueOf(&b))
	utils.DumpObject("i", reflect.ValueOf(&i))
	utils.DumpObject("s", reflect.ValueOf(&s))


	var a []int
	var a1 [1]int
	var a5 [5]int
	var b5 [5]int

	utils.DumpObject("a", reflect.ValueOf(&a))
	utils.DumpObject("a1", reflect.ValueOf(&a1))
	utils.DumpObject("a5", reflect.ValueOf(&a5))
	utils.DumpObject("b5", reflect.ValueOf(&b5))



}




