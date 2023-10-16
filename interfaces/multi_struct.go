package interfaces

import "fmt"

type InterfaceA interface {
	PrintStatus() InterfaceA
}

var interfaceMap = map[string]InterfaceA{
	"A": &A{Status: "nil"},
	"B": &B{Status: "none"},
}

func GetA() *A {
	return interfaceMap["A"].(*A)
}

func GetB() *B {
	return interfaceMap["B"].(*B)
}

type A struct {
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

type B struct {
	Status string
}

func (b *B) ChangeStatus() InterfaceA {
	b.Status = "up"
	return b
}

func (b *B) PrintStatus() InterfaceA {
	fmt.Printf("Status B is %s\n", b.Status)
	return b
}

func StructToInterfaceTest() {
	GetA().ChangeStatus()
	GetA().PrintStatus()
	GetB().ChangeStatus()
	GetB().PrintStatus()
}
