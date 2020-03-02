//Двоичная запись
//Дано натуральное число N. Выведите его представление в двоичном виде.
//
//Входные данные
//Задано единственное число N
//Выходные данные
//Необходимо вывести требуемое представление числа N.
//
//Sample Input:
//12
//Sample Output:
//1100

package main

import . "fmt"

func main() {
	var a int
	Print("Enter a number: ")
	Scan(&a)
	Printf("%b", a)
	// more complex solution
	//sl := make([]int, 0)
	//if a == 0 {
	//	Print("0")
	//} else if a == 1 {
	//	Print("1")
	//} else {
	//	for a >= 1 {
	//		if a % 2 == 0 {
	//			sl = append(sl, 0)
	//		} else {
	//			sl = append(sl, 1)
	//		}
	//		a = a/2
	//	}
	//}
	//for i := len(sl) - 1; i >= 0; i-- {
	//	Print(sl[i])
	//}
}
