package array

import "fmt"

func ArrayCopy() {

	var a = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	editArray(&a)
	fmt.Println(a)
}

func editArray(myArray *[5]int) {

	myArray[3] = 25

	fmt.Println(myArray)

}
