package main

import "fmt"

func main() {

	num := 7

	var ptr = &num

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", num)

}
