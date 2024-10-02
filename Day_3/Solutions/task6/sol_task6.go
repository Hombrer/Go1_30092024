// 3.6 Получить от пользователя натуральное число, посчитать сумму цифр в нём.
//     Решить с помощью индексов, т.е. работаем с числом, как со строкой.

package main

import (
	"fmt"
)

func main() {
	var number string
	var sum int

	fmt.Println("Введите натуральное число:")
	fmt.Scan(&number)

	for idx := range number {
		sum += int(number[idx] - '0') // ASCII  48, val 0
	}

	fmt.Println("Сумма цифр:", sum)
}
