package interfaces

import (
	"fmt"
	"unsafe"
)

type Animal interface {
	Walk()
}
type Dog struct {
	name string
}

func (d Dog) Walk() {
	fmt.Printf("%s is walking \n", d.name)
}

func PointerNilTry() {
	var dog *Dog
	//Test Nils
	nilPointerTest(dog)
	if dog == nil {
		fmt.Println("0")
	} else {
		fmt.Println("1")
	}
	dog = new(Dog)
	//dog = &Dog{}
	//Test Nils
	nilPointerTest(dog)
	if dog == nil {
		fmt.Println("0")
	} else {
		fmt.Println("1")
	}
}

func nilPointerTest(animal Animal) {
	if animal == nil {
		fmt.Println("0")
	} else {
		fmt.Println("1")
	}
}

func NinInterfaceTest() {

	var a Animal

	a.Walk()

}

func Verify() {

	var a Animal
	fmt.Println(a == nil)
	fmt.Printf("a: %T, %v\n", a, a)
	var d *Dog
	fmt.Println(d == nil)
	a = d
	fmt.Println(a == nil)
	fmt.Printf("a: %T, %v\n", a, a)

}

type iface struct {
	itab, data uintptr
}

func Tutorial() {
	var a interface{} = nil
	var b interface{} = (*int)(nil)
	x := 5
	var c interface{} = (*int)(&x)
	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	ic := *(*iface)(unsafe.Pointer(&c))
	fmt.Println(ia, ib, ic)
	fmt.Println(*(*int)(unsafe.Pointer(ic.data)))
}
