package main

import "fmt"

func main() {
	// myarray := [5]int{1, 2, 3, 4, 5}
	// mysubarray := myarray[2:4]
	// fmt.Println(myarray)
	// fmt.Println(mysubarray)
	myMap := make(map[int]string)
	myMap = map[int]string{1: "First", 2: "Second", 3: "third"}
	fmt.Println(myMap)
	myMap[4] = "fouth"
	fmt.Println(myMap)
}
