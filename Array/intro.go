package main

import (
	"fmt"
	"reflect"
)

var (
	name   = "Silas"
	course = "Dockers"
	module = 3.2
	ptr    = &module
)

func main() {
	// name := "Silas"
	// course := "Dockers"
	// module := 3.2
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(module))
	fmt.Println(&module)
	fmt.Println(*ptr)
}
