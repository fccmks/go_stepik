// Дается строка. Нужно удалить все символы, которые встречаются более одного раза
// и вывести получившуюся строку

// Sample Input:
// zaabcbd

// Sample Output:
// zcd

package main

import (
	"fmt"
	"strings"
)

func main() {
	var input, result string
	fmt.Println("Enter a string:")
	fmt.Scan(&input)
	// complicated solution
	// for _, v := range input {
	// 	if !strings.Contains(result, string(v)) {
	// 		result += string(v)
	// 	} else {
	// 		input = strings.Replace(input, string(v), "", -1)
	// 	}
	// }
	// fmt.Println(input)

	// simple solution
	for _, v := range input {
		if strings.Count(input, string(v)) == 1 {
			result += string(v)
		}
	}
	fmt.Println(result)
}
