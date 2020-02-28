//Номер числа Фибоначчи
//Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является,
//то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.
//
//Входные данные
//Вводится натуральное число.
//
//Выходные данные
//Выведите ответ на задачу.
//
//Sample Input:
//8
//Sample Output:
//6

package main

import . "fmt"

func main() {
	var a int
	res := -1
	//Print("Enter a number: ")
	Scan(&a)
	fibslice := make([]int, a)
	fibslice[0] = 1
	fibslice[1] = 1
	for i := 2; i < a; i++ {
		fibslice[i] = fibslice[i-1] + fibslice[i-2]
		if fibslice[i] == a {
			res = i + 1
		}
	}
	Print(res)
}
