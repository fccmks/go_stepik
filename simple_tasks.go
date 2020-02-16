package main

import "fmt"

func main() {
	fmt.Println("Enter numbers: ")
	for n := 0; n <= 100; {
		fmt.Scan(&n)
		if n >= 10 && n <= 100 {
			fmt.Println(n)
		}
	}
}
