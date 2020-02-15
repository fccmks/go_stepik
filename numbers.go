package main

import "fmt"

func main() {
	var a int
	//fmt.Print("Enter a number: ")
	fmt.Scan(&a)
	if a > 0 {
		fmt.Print("Positive number. The next number for the number ", a, " is ", a+1, ".\n")
		fmt.Print("Positive number. The previous number for the number ", a, " is ", a-1, ".")
	} else if a < 0 {
		fmt.Print("Negative number. The next number for the number ", a, " is ", a+1, ".\n")
		fmt.Print("Negative number. The previous number for the number ", a, " is ", a-1, ".")
	}
}
