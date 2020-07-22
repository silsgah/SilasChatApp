package main

import "fmt"

func main() {

	leagueTitle := make(map[string]int)
	leagueTitle["Kwmae"] = 6
	leagueTitle["AKSHD"] = 4

	testMap := map[string]int{"A": 4, "B": 2, "C": 3, "D": 4}

	// for key, value := range testMap {
	// 	fmt.Println("\nKey is : %v Value is: %v", key, value)
	// }
	testMap["A"] = 100
	fmt.Println(testMap)
	delete(testMap, "A")
	fmt.Println(testMap)
}
