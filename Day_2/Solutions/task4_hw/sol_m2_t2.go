/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/
package main

import (
	"fmt"
)

func main() {
	var number int32

	fmt.Println("Введите 4х значное число")
	fmt.Scan(&number)

	first := number/1000
	second := number%1000/100
	third := number%100/10
	last := number%10

	if first == last && second == third {
		fmt.Println("палиндром")
	} else {
		fmt.Println("не палиндром")
	}
}