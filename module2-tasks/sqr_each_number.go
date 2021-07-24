// На вход подается целое число. Необходимо возвести в квадрат каждую цифру
// числа и вывести получившееся число.

// Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81.
// Дальше 1. Единица в квадрате - 1. В итоге получаем 811181

// Sample Input:
// 9119

// Sample Output:
// 811181

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	fmt.Scan(&a)
	s := strconv.Itoa(a)
	for _, v := range s {
		val, _ := strconv.Atoi(string(v))
		fmt.Print(val * val)
	}
}

// подсмотрел на форуме решений
// package main

// import "fmt"

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	for n != 0 {
// 		defer fmt.Print((n % 10) * (n % 10))
// 		n /= 10
// 	}
// }
