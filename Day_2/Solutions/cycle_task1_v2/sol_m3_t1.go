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

import (
	"fmt"
	"strings"
)

func main() {
	var number int64
	fmt.Print("Введите натуральное число: ")
	fmt.Scan(&number)

	var sum int64
	var digits string
	for {
		if number == 0 {
			break 
		}
		var digit int64 = number % 10
		sum += digit
		digits = fmt.Sprintf("%d + ", digit) + digits
		number = number / 10

	}
	// Когда переменная должна быть создана, но она не нужна,
	// можно использовать заглушку в виде _
	digits, _ = strings.CutSuffix(digits, " + ")

	fmt.Println(digits, "=", sum)
}