package interfaces

import (
	"code.byted.org/im_cloud/playground/utils"
	"context"
	"fmt"
	"reflect"
	"strconv"
)

type Person struct {
	name string
	money int
	feature map[string]string
}


type EmptyStruct struct {

}
type EmptyStruct2 struct {

}

func StructCopy()  {

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
	EmptyStruct{} :"afda",
	EmptyStruct2{} : "fffff",
	44:"dahhsdhasd",
}

func EmptyStructPointers (){

	for key, val := range structMap {
		fmt.Printf("key: %+v " , key )
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

func ContextMiddleWare(ctx context.Context, value string)  {

	type myKey struct{}
	ctx = context.WithValue(ctx, myKey{}, value) // Set value
	myValue := ctx.Value(myKey{}).(string)

	fmt.Println(myValue)



a := "dontUseString"
	b := "dontUseString"
	utils.CompareInterfaces(a, b) // true

	// Now let's see how struct{} types behave.
	// Same struct{} keys are always equal:
	utils.CompareInterfaces(StructKey1{}, StructKey1{}) // true
	// Different struct{} keys are never equal even if they appear to be of the same struct{} type:
	utils.CompareInterfaces(StructKey1{}, StructKey2{}) // false

	// This also applies to int keys. Same type means equal:
	utils.CompareInterfaces(IntKey1(0), IntKey1(0)) // true
	// And different keys are never equal even though they have 0 value:
	utils.CompareInterfaces(IntKey1(0), IntKey2(0)) // false
	// However, unlike struct{}, an int typed key allows for mistakes with the value:
	utils.CompareInterfaces(IntKey1(0), IntKey1(1)) // false
	// To add to why you shouldn't use int typed keys, when seeing that the package define a
	// key as int, the user might be tempted to pass a primitive 0 instead. Which doesn't work:
	utils.CompareInterfaces(IntKey1(0), 0) // false

}