//Количество минимумов
//Найдите количество минимальных элементов в последовательности.
//
//Входные данные
//Вводится натуральное число N, а затем N чисел.
//
//Выходные данные
//Выведите количество минимальных элементов.
//
//Sample Input:
//3
//21
//11
//4
//Sample Output:
//1

package main

import . "fmt"

func main() {
	var a, count int
	Print("Numbers qnty: ")
	Scan(&a)
	nums := make([]int, a)
	Print("Enter slice elements: ")
	for i := 0; i < a; i++ {
		Scan(&nums[i])
	}
	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
			count = 1
		} else if v == min {
			count++
		}
	}
	Print(count)
}
