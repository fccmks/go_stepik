package main

import "fmt"

//Программа должна вывести цифры, которые имеются в обоих числах, через пробел.
//Цифры выводятся в порядке их нахождения в первом числе.
//Sample Input:
//564 8954
//Sample Output:
//5 4
func main() {
	var a, b int
	kForA := 10
	fmt.Print("Enter numbers: ")
	fmt.Scan(&a, &b)
	if a > 100 && a < 1000 {
		kForA = 100
	} else if a > 1000 {
		kForA = 1000
	}
	for kForA > 0 {
		etalon := a / kForA
		if etalon > 10 {
			etalon = etalon % 10
		}
		defB := b
		for defB > 0 {
			eq := defB % 10
			if etalon == eq {
				fmt.Print(etalon, " ")
			}
			defB = defB / 10
		}
		kForA = kForA / 10
	}
}
