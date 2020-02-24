//По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова",
//"n коровы", правильно склоняя слово "корова".
//
//Входные данные
//Дано число n (0<n<100).
//
//Выходные данные
//Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.
//
//Sample Input:
//10
//Sample Output:
//10 korov

package main

import . "fmt"

func main() {
	var a int
	word := "korov"
	Print("Enter cows quantity: ")
	Scan(&a)
	if a < 10 || a > 20 {
		if a%10 == 1 {
			word = word + "a"
		} else if a%10 > 1 && a%10 < 5 {
			word = word + "y"
		}
	}
	Println(a, word)
}
