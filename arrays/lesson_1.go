//Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4),
//а затем NN целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
//Sample Input:
//5
//41 -231 24 49 6
//Sample Output:
//49
package main

import . "fmt"

func main() {
	var a int
	Print("Enter slice length: ")
	Scan(&a)
	nums := make([]int, a)
	Print("Enter slice elements: ")
	for i := 0; i < a; i++ {
		Scan(&nums[i])
	}
	Println(nums[3])
}