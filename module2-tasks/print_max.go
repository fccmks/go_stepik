// Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.

// Входные данные
// Вводится строка ненулевой длины. Известно также, что длина строки
// не превышает 1000 знаков и строка содержит только десятичные цифры.

// Выходные данные
// Выведите максимальную цифру, которая встречается во введенной строке.

// Sample Input:
// 1112221112

// Sample Output:
// 2

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	max := 0
	for _, v := range s {
		var a int
		a, _ = strconv.Atoi(string(v))
		if a > max {
			max = a
		}
	}

	// var s []byte
	// fmt.Scan(&s)
	// for i, _ := range s {
	// 	fmt.Println(string(s[i]))
	// }

	fmt.Println(max)
}
