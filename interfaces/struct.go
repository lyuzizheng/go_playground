package interfaces

import (
	"code.byted.org/lyuzizheng/playground/utils"
	"fmt"
	"reflect"
	"strconv"
)

type Person struct {
	name    string
	money   int
	feature map[string]string
}

type EmptyStruct struct {
}
type EmptyStruct2 struct {
}

func StructCopy() {

	person1 := Person{
		name:    "person1",
		money:   100,
		feature: make(map[string]string),
	}

	person1.feature["age"] = "jaja"
	person1.feature["personality"] = "good"

	for i := 0; i < 20; i++ {
		person2 := person1
		person2.feature = make(map[string]string)
		person2.feature[strconv.Itoa(i)] = strconv.Itoa(i)

		fmt.Printf("Person2: %+v \n", person2)
	}

}

var structMap = map[interface{}]string{
	EmptyStruct{}:  "afda",
	EmptyStruct2{}: "fffff",
	44:             "dahhsdhasd",
}

func EmptyStructPointers() {

	for key, val := range structMap {
		fmt.Printf("key: %+v ", key)
		fmt.Println("val: ", val)

		utils.PrintObjectDumpTableHeader()
		utils.DumpObject("key", reflect.ValueOf(&key))
		utils.DumpObject("val", reflect.ValueOf(&val))

	}
	//
	//utils.PrintObjectDumpTableHeader()
	//utils.DumpObject("Empty Struct map", reflect.ValueOf(&structMap))

	a := structMap[EmptyStruct{}]
	b := structMap[EmptyStruct2{}]

	fmt.Println(a)
	fmt.Println(b)

	utils.PrintObjectDumpTableHeader()
	utils.DumpObject("key", reflect.ValueOf(&a))
	utils.DumpObject("val", reflect.ValueOf(&b))

}

type StructKey1 struct{}
type StructKey2 struct{}

type IntKey1 int
type IntKey2 int

func InterfaceCompare() {

	//type myKey struct{}
	//ctx = context.WithValue(ctx, myKey{}, value) // Set value
	//myValue := ctx.Value(myKey{}).(string)

	a := 1

	b := 1

	CompareInterfaces(StructKey1{}, StructKey1{})
	CompareInterfaces(StructKey1{}, StructKey2{})
	CompareInterfaces(IntKey1(0), IntKey2(0))
	CompareInterfaces(IntKey1(0), 0)
	CompareInterfaces(a, 1)
	CompareInterfaces(IntKey1(a), IntKey1(b))

}

func CompareInterfaces(a interface{}, b interface{}) {
	fmt.Println(a == b)
}
