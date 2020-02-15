package main

import "fmt"

func main() {
	var x, y, p float32
	counter := 0
	fmt.Print("Enter: ")
	fmt.Scan(&x, &p, &y)
	for x < y {
		x = x + x/100*p
		fmt.Printf("it is %d year. sum is %v\n", counter, x)
		counter++
	}
	fmt.Println(counter)
}
