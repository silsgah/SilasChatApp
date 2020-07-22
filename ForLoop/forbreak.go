package main

import (
	"time"

	"github.com/mattn/go/src/fmt"
)

func main() {
	for timer := 10; timer >= 0; timer-- {
		if timer%2 == 0 {
			continue
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}
}
