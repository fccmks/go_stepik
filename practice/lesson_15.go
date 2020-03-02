//Из натурального числа удалить заданную цифру.
//
//Входные данные
//Вводятся натуральное число и цифра, которую нужно удалить.
//Выходные данные
//
//Вывести число без заданных цифр.
//
//Sample Input:
//38012732
//3
//Sample Output:
//801272

package main

import . "fmt"

func main() {
	var a, rem int
	var sl []int
	Print("Enter numbers: ")
	Scan(&a, &rem)
	for a > 0 {
		if a%10 != rem {
			sl = append(sl, a%10)
		}
		a /= 10
	}
	for i := len(sl) - 1; i >= 0; i-- {
		Print(sl[i])
	}
}
