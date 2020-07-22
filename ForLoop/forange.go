package main

import "github.com/mattn/go/src/fmt"

func main() {

	courseInProg := []string{"Docker Deep Dive",
		"Docker Clustering", "Docker and Kubernetes"}
	courseCompleted := []string{"Docker Deep Dive", "Go completed",
		"Go fundamental completed"}
	for _, i := range courseInProg {
		fmt.Println(i)
		for _, j := range courseCompleted {
			if i == j {
				fmt.Println("\nHey we found a class", i, "in both list")
			}
		}
	}
}
