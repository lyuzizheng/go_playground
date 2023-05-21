package main

import (
	"code.byted.org/lyuzizheng/playground/sug_reply"
	"fmt"
)

func main() {

	sug_reply.SugReply()

}

type InterfaceA interface {
	PrintStatus() InterfaceA
	ChangeStatus() InterfaceA
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

type TestStruct struct {
	*A
}

func (t *TestStruct) ChangeStatus() InterfaceA {
	t.Status = "up"
	return t
}

func (t *TestStruct) PrintStatus() InterfaceA {
	t.A.PrintStatus()
	//fmt.Printf("Status T is %s\n", t.Status)
	return t
}
