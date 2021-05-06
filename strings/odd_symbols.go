// На вход дается строка, из нее нужно сделать другую строку, оставив только
// нечетные символы (считая с нуля)

// Sample Input:
// ihgewlqlkot

// Sample Output:
// hello

package main

import (
	"fmt"
)

func main() {
	var text string
	fmt.Println("Write a string:")
	fmt.Scan(&text)
	for k, v := range text {
		if k%2 != 0 {
			fmt.Print(string(v))
		}
	}
}
