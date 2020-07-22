package main

import (
	"fmt"
)

func main() {
	// var waitGrp sync.WaitGroup
	// waitGrp.Add(2)
	go func() {
		// defer waitGrp.Done()
		//time.Sleep(5 * time.Second)
		fmt.Println("Hello from func 1")
	}()

	go func() {
		// defer waitGrp.Done()
		fmt.Println("hello from func 2")
	}()
	// time.Sleep(10 * time.Second)
	fmt.Println("Hell from main goroutine")
	// waitGrp.Wait()
}
