/*
Задача №1
Вводим любое натуральное число.
Нужно посчитать сумму цифр числа, с помощью цикла for

Пример:
In: 4521
Out: 12
*Задание 1.1: 4 + 5 + 2 + 1 = 12 - добавить к выводу сумму как выражение
*/
package main

import "fmt"

func main() {
	var n, sum int
	var statement string

	fmt.Println("Введите число: ")
	fmt.Scan(&n)

	for n != 0 {
		digit := n % 10
		if len(statement) == 0 {
			statement = fmt.Sprintf("%d = ", digit)
		} else {
			statement = fmt.Sprintf("%d + ", digit) + statement
		}
		sum += digit
		n /= 10
	}

	fmt.Printf("Сумма числа: %v\n", sum)
	fmt.Printf("%s%v\n", statement, sum)
}
