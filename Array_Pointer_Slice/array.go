package main

import "fmt"

func main() {
	var subjects [5]string

	subjects[0] = "AFLL"
	subjects[1] = "DDCO"
	subjects[2] = "WT"
	subjects[3] = "MCSE"
	subjects[4] = "DSA"

	fmt.Println("The sub for 3rd sem are: ", subjects)
	fmt.Println("The num of sub are", len(subjects))

	var subjects2 = [1]string{"CIE"}
	fmt.Println("No of 2 credit course is", len(subjects2))

}
