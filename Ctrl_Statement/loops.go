package main

import (
	"fmt"
)

func main() {
	fmt.Println("In this code we will be discussing about loops \n. In GoLang there is only for loop and it takes the form of other loop")

	days := []string{"Mon", "Tue", "Wed", "Thur", "Fri", "Sat", "Sun"}
	fmt.Println(days)

	//TYPE 1
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//TYPE 2
	for i := range days {
		fmt.Println(days[i])
	}

	//TYPE 3
	i := 0
	for i < len(days) {
		fmt.Println(days[i])
		i++
	}
}
