//На ввод подаются пять чисел, которые записываются в массив.
//Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.
//
//Sample Input:
//2
//3
//56
//45
//21
//Sample Output:
//56

package main

import . "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		Scan(&a)
		array[i] = a
	}
	max := array[0]
	for _, v := range array {
		if v > max {
			max = v
		}
	}
	Println(max)
}
