package main

import "github.com/mattn/go/src/fmt"

func main() {
	name := "Silas"
	course := "Docker Deep Dive"
	const price = 748
	fmt.Println("\nHi", name, " you are currently watching", course)
	changeCourse(&course)

	fmt.Println("\nNew course is", course, price)
}

func changeCourse(course *string) string {
	*course = "First look: Natove Docker Clustering"
	fmt.Println("\nTrying to change your course to", *course)
	return *course
}

//Variable by Value
// func main() {
// 	name := "Silas"
// 	course := "Docker Deep Dive"

// 	fmt.Println("\nHi", name, " you are currently watching", course)
// 	changeCourse(course)

// 	fmt.Println("\nNew course is", course)
// }

// func changeCourse(course string) string {
// 	course = "First look: Natove Docker Clustering"
// 	fmt.Println("\nTrying to change your course to", course)
// 	return course
// }
