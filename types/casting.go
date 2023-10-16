package types

import (
	"fmt"
	"playground/utils"
	"reflect"
	"unsafe"
)

func TypeCasting() {

	var a = make([]byte, 0, 15)
	var b = "hello"
	a = append(a, []byte(b)...)

	var c = *(*string)(unsafe.Pointer(&a))

	fmt.Println(string(a))
	fmt.Println(c)
	a = append(a, []byte(b)...)

	a[0] = 'k'
	a[7] = '2'

	fmt.Println(string(a))
	fmt.Println(c)

}

func TypeCasting2() {

	var a = make([]byte, 0, 15)
	var b = "hello"
	a = append(a, []byte(b)...)

	var c = *(*string)(unsafe.Pointer(&a))

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(&c)
	utils.PrintObjectDumpTableHeader()
	utils.DumpObject("c", reflect.ValueOf(&a))
	utils.DumpObject("c", reflect.ValueOf(&c))

	fmt.Println(len(a))
	fmt.Println(len(c))

	a = append(a, []byte(b)...)

	a[0] = 'k'
	a[7] = '2'

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(&c)
	utils.PrintObjectDumpTableHeader()
	utils.DumpObject("c", reflect.ValueOf(&a))
	utils.DumpObject("c", reflect.ValueOf(&c))

	fmt.Println(len(a))
	fmt.Println(len(c))

}
