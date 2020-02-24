//Самое большое число, кратное 7
//Найдите самое большее число на отрезке от a до b, кратное 7 .
//
//Входные данные
//Вводится два целых числа a и b (a≤b).
//
//Выходные данные
//Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b), кратное 7 ,
//или выведите "NO" - если таковых нет.
//
//Sample Input:
//100
//500
//Sample Output:
//497

package main

import . "fmt"

func main() {
	var a, b, max int
	Print("Enter two numbers: ")
	Scan(&a, &b)
	for i := a; i <= b; i++ {
		if i%7 == 0 && i > max {
			max = i
		}
	}
	if max == 0 {
		Println("NO")
	} else {
		Println(max)
	}
}
