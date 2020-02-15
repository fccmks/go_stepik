package main

import fmt "fmt"

func main() {
	var a int
	fmt.Print("Enter a number: ")
	fmt.Scan(&a)
	if (a%4 == 0 && a%100 != 0) || a%400 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
