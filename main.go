package main

import (
	"code.byted.org/im_cloud/playground/interface_test"
	"fmt"
	"net/http"
)

var a = &http.Client{}
var b = &http.Client{}


func main() {

	fmt.Printf("%p : a pointer\n", a)
	fmt.Printf("%p : b pointer\n", b)

	a := interface_test.AutoScalerServiceImpl{}
	a.ParseMethod("lalalal").ParseParameters("commandInlineString").Redirect()




}