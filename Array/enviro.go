package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USERNAME")
	fmt.Println(name)
	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }
}
