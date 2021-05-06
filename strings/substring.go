// Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X.
// Если подстроки S нет в строке X - вывести -1

// Sample Input:
// awesome
// es

// Sample Output:
// 2

package main

import (
	"fmt"
	"strings"
)

func main() {
	var text1, text2 string
	fmt.Println("Write a string:")
	// text1, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// text2, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// text1 = strings.Trim(text1, "\n")
	// text2 = strings.Trim(text2, "\n")
	fmt.Scan(&text1)
	fmt.Scan(&text2)
	fmt.Println(strings.Index(text1, text2))
}
