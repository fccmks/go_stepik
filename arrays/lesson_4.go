//Дан массив, состоящий из целых чисел. Напишите программу,
//которая подсчитывает количество положительных чисел среди элементов массива.
//
//Входные данные
//Сначала задано число NN — количество элементов в массиве (1\leq N\leq1001≤N≤100).
//Далее через пробел записаны NN чисел — элементы массива. Массив состоит из целых чисел.
//
//Выходные данные
//Необходимо вывести единственное число - количество положительных элементов в массиве.
//
//Sample Input:
//5
//1 2 3 -1 -4
//Sample Output:
//3

package main

import . "fmt"

func main() {
	var a, count int
	Print("Enter slice length: ")
	Scan(&a)
	nums := make([]int, a)
	Print("Enter slice elements: ")
	for i := 0; i < a; i++ {
		Scan(&nums[i])
	}
	for _, v := range nums {
		if v > 0 {
			count++
		}
	}
	Print(count)
}