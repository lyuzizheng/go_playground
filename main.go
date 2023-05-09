package main

import "fmt"

func main() {

	s1 := make([]int, 3, 6)
	fmt.Println(s1)
	s2 := s1
	fmt.Println(s2)
	s1 = append(s1, 7)
	fmt.Println(s1)
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(s1)

	//array.ArrayCopy()
	//interfaces.PointerNilTry()
	//interfaces.Verify()
	//interfaces.Tutorial()
	//channel.ChannelTest1()   //C
	//system.ChangeTime()

	//interfaces.EmptyStructPointers()
	//algo.RunGame()
	//uid := int64(6691213328788014082)
	//
	//fmt.Println(uid)
	//
	//remainder := uid & 0xFFF
	//
	//fmt.Println(remainder)

	//context.TestContextCancel()
	//interfaces.InterfaceCompare()

	//json.UnmarshallToInterface()
	//types.TypeCasting()
	//json.JsonEncode()

	//pointers.PointerDemo()

	//interfaces.CheckSliceSimple()

	//var s interfaces{}
	//fmt.Println(unsafe.Sizeof(s))

	//utils.DumpObject("interfaces", reflect.ValueOf(&s))

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

func updateSlice(s1 []int) {
	s1 = append(s1, 5, 5, 5, 5, 5)
}
