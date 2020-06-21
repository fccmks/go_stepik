//Напишите функцию sumInt, принимающую переменное количество аргументов типа int,
//и возвращающую количество полученных функцией аргументов и их сумму.
//
//Пример вызова вашей функции:
//
//a, b := sumInt(1, 0)
//fmt.Println(a, b)
//
//Результат: 2, 1

package main

import "fmt"

func main() {
	fmt.Print(sumInt(1, 0))
}

func sumInt(n ...int) (argscount, argssum int) {
	argssum = 0
	for _, v := range n {
		argssum += v
	}
	return len(n), argssum
}
