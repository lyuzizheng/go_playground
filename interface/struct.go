package _interface

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	money int
	feature map[string]string
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