package main

import "fmt"

func main() {
	fmt.Println("In this code, we will discuss if-else statements.")

	marks := 80

	if marks >= 90 && marks <= 100 {
		fmt.Println("A grade")
	} else if marks >= 80 && marks < 90 {
		fmt.Println("B grade")
	} else if marks >= 60 && marks < 80 {
		fmt.Println("C grade")
	} else {
		fmt.Println("Fail")
	}

	fmt.Printf("Your marks are: %d\n", marks)
}
