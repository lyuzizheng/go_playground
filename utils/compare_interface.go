package utils

import (
	"fmt"
	"unsafe"
)

func CompareInterfaces(key1 interface{}, key2 interface{}) {
	type ifaceHdr struct {
		T unsafe.Pointer
		V unsafe.Pointer
	}

	fmt.Println("\ninterface1 == interface2?", key1 == key2)
	fmt.Printf("interface1 %+v\n", *(*ifaceHdr)(unsafe.Pointer(&key1)))
	fmt.Printf("interface2 %+v\n", *(*ifaceHdr)(unsafe.Pointer(&key2)))
}
