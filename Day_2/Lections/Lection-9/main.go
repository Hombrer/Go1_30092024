package main

import (
	"fmt"
	"math"
)

func main()  {
	// Массивы. Основа
	// 1. Определение массива
	// Создание массива из 5-ти целочисленных значений
	// При инициализации массива важно передать информацию - сколько элементов в нем будет
	var arr [5]int 
	fmt.Println("This is my array:", arr)

	// 2. Присваивание значения элементу массиву
	// Необходимо обратиться к элементу массива через его порядковый номер(индекс):
	// arr[i] = elem
	arr[0] = 10
	arr[1] = 20
	arr[3] = -100
	arr[4] = 1000
	fmt.Println("Array after change:", arr)

	// 3. Определение массива с указанием элементов на месте
	// Если при инициализации кол-во элементов меньше номинальной длины, 
	// то недостающие элементы инициализируются нулями
	newArr := [5]int{10, 20, 30}
	fmt.Println("Short declaration and init:", newArr)

	// 4. Создание массива через инициализацию переменных
	arrWithValues := [...]int{10, 20, 30, 40}
	fmt.Println("Array declaration with [...]:", arrWithValues, len(arrWithValues))
	arrWithValues[0] = 10_000
	fmt.Println("Array declaration with [...]:", arrWithValues, len(arrWithValues))

	// 5. Массив - это набор ЗНАЧЕНИЙ. То есть при всех манипуляциях - массив копируется
	// (жестко, на уровне компилятора)
	first := [...]int{1, 2, 3}
	second := first
	second[0] = 1000
	fmt.Println("first:", first)
	fmt.Println("second:", second)

	// 6. Массив и его размер - это две составляющие одного типа (размер массива - это часть типа)
	// var third [2]int
	// third = first  // Error

	// 7. Итерирование по массиву
	floatArray := [...]float64{12.5, 10.0, 12.1, 15.2, 13.51}
	for i := 0; i < len(floatArray); i++ {
		fmt.Printf("%d element of array is %.2f\n", i + 1, floatArray[i])
	}

	fmt.Println("-- Repeat --")
	// 8. Итерирование по массиву через оператор range
	var sum float64
	for idx, value := range floatArray {
		fmt.Printf("%d element of array is %.2f\n", idx + 1, floatArray[idx])
		sum += value
	}
	fmt.Println("Total sum:", sum)


	// 9. Игнорирование idx при обращении к массиву через оператор range
	var summa float64
	for _, value := range floatArray {
		summa += value
	}
	fmt.Println("Repeat total sum:", math.Round(summa * 100) / 100)
	fmt.Println("Repeat total sum:", math.Round(12.5))
	fmt.Println("Repeat total sum:", math.Round(11.5))

	// 10. Многомерные массивы
	words := [2][2]string{
		{"Mark", "Alice"},
		{"Viktor", "Inna"},
	}
	fmt.Println("Multidimensional array:", words)
	// 11. Итерирование по многомерному массиву
	for _, innerArr := range words {
		for _, word := range innerArr {
			fmt.Printf("Word: %s ", word)
		}
		fmt.Println()
	}
}