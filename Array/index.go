package main

import "fmt"

func main() {
	myarray := [5]int{1, 2, 3, 4, 5}
	mysubArray := myarray[2:4]
	// mysubArray[0] = 2
	fmt.Println(myarray)
	fmt.Println(mysubArray)
}
