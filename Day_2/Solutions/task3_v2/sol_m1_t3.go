package main

import (
	"fmt"
	"math"
)

/*
Задача №3

На входе размер вклада(float64), кол-во лет(float64) и процент по вкладу(int64)
Проверить условия (от и до включительно):
вклад от 100 до 1_000_000
кол-во лет от 1 до 100
процент от 1 до 20

и посчитать размер вклада при выполнении условий.
В противном случае вывести сообщение о неправильных данных

Использовать ежегодную капитализацию.
*/

func main() {
	var depositAmount float64
	fmt.Println("Введите размер вклада")
	fmt.Scan(&depositAmount)

	if depositAmount < 100 || depositAmount > 1_000_000 {
		fmt.Println("Размер вклада должен быть от 100 до 1_000_000.")
		return
	}

	var depositYears int64
	fmt.Println("Введите кол-во лет вклада")
	fmt.Scan(&depositYears)

	if depositYears < 1 || depositYears > 100 {
		fmt.Println("Кол-во лет вклада должно быть от 1 до 100.")
		return
	}

	var depositPercents int64
	fmt.Println("Введите проценты по вкладу")
	fmt.Scan(&depositPercents)

	if depositPercents < 1 || depositPercents > 20 {
		fmt.Println("Проценты по вкладу должны быть от 1 до 20.")
		return
	}

	var result = depositAmount * math.Pow(1+float64(depositPercents)/100, float64(depositYears))
	fmt.Printf("Итоговый размер вклада c начальной суммой %.2f под %d%% на %d лет: %.2f\n",
		depositAmount, depositPercents, depositYears, result)
}
