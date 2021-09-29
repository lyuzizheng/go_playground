package interface_test

import (
	"fmt"
)


//go:noinline
func UpdateSlice(s []int)  {

	s[2] = 0

	fmt.Printf("Slice2: %v\n", s)

	s = append(s, 8)

	fmt.Printf("Slice3: %v\n", s)

	s[1] = 5

	fmt.Printf("Slice4: %v\n", s)
}


func CheckSlice(){

	var s = make([]int, 3, 10)
	s[0] = 1
	s[1] = 2
	s[2] = 3

	//var s = []int{1,2,3}

	fmt.Printf("Slice1: %v\n", s)

	UpdateSlice(s)

	fmt.Printf("Slice5: %v\n", s)
}

func CheckSliceSimple() {

	var s = make([]int, 3, 5)  // -> 24bytes

	//{
	//	unsafe.pointer ->Actual Array (never changes)
	//  len -> current leng
	//  capacity ->
	//}
	s[0] = 10
	s[1] = 9
	s[2] = 8 // s = 10, 9, 8

	var d = s // Using same underneath array with s (but they are different struct)

	s = append(d, 7,6,5) // s = 10, 9, 8, 7, 6, 5  len(s) => 6 and len(d) is still 3
	d[0] = 4 //s[0] = d[0] 4
	s[4] = 3 // s[4] = 3 s = 4 9 8 7 3 5
	d = append(d, 2)  //d and s sharing same array len(d) = 3  s[3] =d[3] = 2

	fmt.Printf("SliceS: %v\n", s)
	fmt.Printf("SliceD: %v\n", d)

}
