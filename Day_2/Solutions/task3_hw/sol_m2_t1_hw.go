/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1 9 2
Выход: 1 2 9

Про sort мы пока ничего не знаем!
Это задача на условный оператор
*/
package main

import (
	"fmt"
)

func main() {

	var (
		a, b, c int16
	)

	fmt.Println("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Println("Введите второе число: ")
	fmt.Scan(&b)
	fmt.Println("Введите третье число: ")
	fmt.Scan(&c)
	fmt.Printf("Ввод: %v %v %v\n", a, b, c)

	if a > b {
		a, b = b, a
	}

	if a > c {
		a, c = c, a
	}

	if b > c {
		b, c = c, b
	}

	fmt.Printf("Вывод: %v %v %v\n", a, b, c)
}

/*
	9 6 3
	9 3 6
	6 3 9
	3 6 9
*/
