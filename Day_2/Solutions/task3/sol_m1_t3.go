/* 
Задача №3

На входе размер вклада(float64), кол-во лет(int64) и процент по вкладу(int64)
Проверить условия (от и до включительно): 
вклад от 100 до 1_000_000
кол-во лет от 1 до 100
процент от 1 до 20

и посчитать размер вклада при выполнении условий.
В противном случае вывести сообщение о неправильных данных.

Использовать ежегодную капитализацию.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		deposit float64
		number_of_years int64
		deposit_percent int64
	)
	fmt.Printf("Введите размер вклада, кол-во лет и процент по вкладу:\n")
	fmt.Scan(&deposit, &number_of_years, &deposit_percent)
	if (100 <= deposit && deposit <= 1_000_000) && (1 <= number_of_years && number_of_years <= 100) && (1 <= deposit_percent && deposit_percent <= 20) {
		s := deposit * math.Pow(1+((float64(deposit_percent)/100)), float64(number_of_years))

		fmt.Printf("Размер вклада = %.2f рублей \n", s)
	} else {
		fmt.Println("Данные не корректны")
	}
}