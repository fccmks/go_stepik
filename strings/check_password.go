// Ваша задача сделать проверку подходит ли пароль вводимый пользователем
// под заданные требования. Длина пароля должна быть не менее 5 символов,
// он должен содержать только цифры и буквы латинского алфавита.
// На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok",
// иначе вывести "Wrong password"

// Sample Input:
// fdsghdfgjsdDD1

// Sample Output:
// Ok

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var input string
	minPswLength := 5
	fmt.Println("Enter a string:")
	fmt.Scan(&input)
	if utf8.RuneCountInString(input) < minPswLength {
		fmt.Println("Wrong password")
	} else {
		rs := []rune(input)
		for i := range rs {
			if unicode.Is(unicode.Latin, rs[i]) || unicode.IsDigit(rs[i]) {
				continue
			} else {
				fmt.Println("Wrong password")
				return
			}
		}
		fmt.Println("Ok")
	}
}
