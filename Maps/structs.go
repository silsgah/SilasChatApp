package main

import "github.com/mattn/go/src/fmt"

func main() {

	type CourseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	// var DockerDeepDive CourseMeta
	// DockerDeepDive := new(CourseMeta)
	DockerDeepDive := CourseMeta{
		Author: "Silas Haj",
		Level:  "Advanced",
		Rating: 5,
	}

	fmt.Println(DockerDeepDive)
	fmt.Println(DockerDeepDive.Author)
}
