//Дано трехзначное число. Найдите сумму его цифр.
//
//Формат входных данных
//На вход дается трехзначное число.
//
//Формат выходных данных
//Выведите одно целое число - сумму цифр введенного числа.
//
//Sample Input:
//745
//Sample Output:
//16

package main

import . "fmt"

func main() {
	var a, sum int
	//Print("Enter number: ")
	Scan(&a)
	for a > 0 {
		sum += a % 10
		a = a / 10
	}
	Print(sum)
}
