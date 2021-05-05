// На вход подается строка. Нужно определить, является ли она палиндромом.
// Если строка является паллиндромом - вывести Палиндром иначе - вывести Нет.
// Все входные строчки в нижнем регистре.

// Sample Input:
// топот

// Sample Output:
// Палиндром

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Write a string:")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")
	textLen := len([]rune(text))
	for i := 0; i < textLen/2; i++ {
		letterFromStart := string([]rune(text)[i])
		letterFromEnd := string([]rune(text)[textLen-i-1])
		if letterFromStart != letterFromEnd {
			fmt.Println("Нет")
			return
		}
	}
	fmt.Println("Палиндром")
}
