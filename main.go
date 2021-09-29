package main

import (
	"code.byted.org/im_cloud/playground/interface_test"



	"fmt"
)

func main() {


	//interface_test.PointerNilTry()
	//interface_test.Verify()
	interface_test.Tutorial()



	//jsonPlay.UnmarshallToInterface()
	//types.TypeCasting()
	//jsonPlay.JsonEncode()

	//pointers.PointerDemo()

	//interface_test.CheckSliceSimple()

	//var s interface{}
	//fmt.Println(unsafe.Sizeof(s))

	//utils.DumpObject("interface", reflect.ValueOf(&s))

	//var testStruct InterfaceA

	//testStruct = &TestStruct{
	//	A: &A{Status: "Init"},
	//}

	//testStruct.PrintStatus().ChangeStatus().PrintStatus()

	//type S struct {
	//	A uint8
	//	B uint8
	//	C uint64
	//	D uint64
	//}
	//
	//type S1 struct {
	//	A uint16
	//	B uint32
	//	C uint64
	//	D uint64
	//}
	//
	//type S2 struct {
	//	A uint8
	//	B uint64
	//	C uint64
	//	D uint64
	//}
	//
	//type S3 struct {
	//	A uint32
	//	B uint32
	//	C uint64
	//	D uint64
	//}
	//
	//
	//
	//fmt.Println(unsafe.Sizeof(S{}))
	//fmt.Println(unsafe.Sizeof(S1{}))
	//fmt.Println(unsafe.Sizeof(S2{}))
	//fmt.Println(unsafe.Sizeof(S3{}))



	//a := make([]int8, 1024)
	//b := make([]int8, 1024)
	//fmt.Printf("a: %p, b: %p\n", a, b)
	//fmt.Printf("a: %p, b: %p\n", &a[0], &a[1])
	//fmt.Printf("a: %p, b: %p\n", &a[1023], &b[0])

}


type InterfaceA interface {
	PrintStatus() InterfaceA
	ChangeStatus() InterfaceA
}


type A struct{
	Status string
}

func (a *A) ChangeStatus() InterfaceA {
	a.Status = "down"
	return a
}

func (a *A) PrintStatus() InterfaceA {
	fmt.Printf("Status A is %s\n", a.Status)
	return a
}


type TestStruct struct {
	*A
}

func (t *TestStruct) ChangeStatus() InterfaceA{
	t.Status = "up"
	return t
}

func (t *TestStruct) PrintStatus() InterfaceA{
	t.A.PrintStatus()
	//fmt.Printf("Status T is %s\n", t.Status)
	return t
}
