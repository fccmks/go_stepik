// На вход подается строка. Нужно определить, является ли она правильной или нет.
// Правильная строка начинается с заглавной буквы и заканчивается точкой.
// Если строка правильная - вывести Right иначе - вывести Wrong

// Маленькая подсказка: fmt.Scan() считывает строку до первого пробела, вы можете считать
// полностью строку за раз с помощью bufio:

// text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

// Sample Input:
// Быть или не быть.

// Sample Output:
// Right

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Write a string:")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")
	textLen := len([]rune(text))
	if unicode.IsUpper([]rune(text)[0]) && []rune(text)[textLen-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
