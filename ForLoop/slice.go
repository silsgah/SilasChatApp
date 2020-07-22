package main

import "fmt"

func main() {
	// mycourse := make([]string, 5, 10)

	// fmt.Println("Length is: %d.\nCapacity is: %d",
	// 	len(mycourse), cap(mycourse))
	// mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(mySlice[4])
	// mySlice[1] = 0
	// fmt.Println(mySlice)

	// sliceOfSlice := mySlice[2:5]
	// fmt.Println(sliceOfSlice)

	mySlices := make([]int, 1, 4)

	fmt.Println("length is: %d Capacity is : %d",
		len(mySlices), cap(mySlices))
	for i := 1; i < 17; i++ {
		mySlices = append(mySlices, i)
		fmt.Println("\nCapacity is : %d", cap(mySlices))
	}
}
