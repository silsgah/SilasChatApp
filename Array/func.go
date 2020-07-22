package main

import (
	"fmt"
	"strings"
)

func main() {
	// module := "function basics"
	// author := "Silas Gah"
	bestFinish := bestGrade(12, 34, 213, 45, 45, 45, 45)
	fmt.Println(bestFinish)
	// fmt.Println(converter(module, author))
}
func bestGrade(grade ...int) int {
	best := grade[0]
	for _, i := range grade {
		if i < best {
			best = i
		}
	}
	return best
}
func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)

	return module, author
}
