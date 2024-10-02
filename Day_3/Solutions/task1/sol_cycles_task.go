package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var a, b float64
	fmt.Println("Введите два числа a и b через пробел, такие, что b > a.")
	fmt.Scan(&a, &b)

	for a >= b {
		fmt.Println("Число a должно быть больше b. Попробуйте еще раз.")
		fmt.Scan(&a, &b)
	}

	var intA = int(a)
	intB := int(b)
	if math.Mod(b, 1) > 0 {
		intB++
	}

	if b-a <= 1 {
		fmt.Printf("Между %.2f и %.2f нет целых чисел.\n", a, b)
	} else {
		fmt.Printf("Все целые числа между %.2f и %.2f:\n", a, b)
		for i := intA + 1; i < intB; i++ {
			fmt.Println(i)
		}
	}

	firstDivisibleByFive := (intA/5 + 1) * 5
	if firstDivisibleByFive > intB {
		fmt.Printf("Между %.2f и %.2f нет чисел, делящихся на 5 без остатка.\n", a, b)
	} else {
		fmt.Printf("Все числа между %.2f и %.2f, делящиеся на 5 без остатка:\n", a, b)
		for i := firstDivisibleByFive; i < intB; i += 5 {
			fmt.Println(i)
		}
	}

	var number int
	fmt.Println("Введите число для получения таблицы умножения.")
	fmt.Scan(&number)
	fmt.Println("Ваша таблица умножения:")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}

	var sum uint
	for i := 1; i <= 4; i++ {
		fmt.Println("Введите оценку по экзамену №", i)
		var score uint
		fmt.Scan(&score)
		sum += score
	}
	fmt.Printf("Сумма набранных баллов: %d\n", sum)

	fmt.Println("Теперь вы можете неограниченно вводить числа для суммирования.")
	fmt.Println("Как только надоест, введите -1 для остановки.")
	var inputNumber int
	var sumOfInputs int
	for inputNumber != -1 {
		sumOfInputs += inputNumber
		fmt.Scan(&inputNumber)
	}
	fmt.Printf("Сумма ваших чисел: %d\n", sumOfInputs)

	fmt.Println("Введите натуральное число для подсчета суммы цифр")
	var naturalNumber uint
	fmt.Scan(&naturalNumber)
	var strNaturalNumber = fmt.Sprintf("%d", naturalNumber)
	var sumOfNumbers int
	for i := 0; i < len(strNaturalNumber); i++ {
		ch := fmt.Sprintf("%c", strNaturalNumber[i])
		var res, _ = strconv.Atoi(ch)
		sumOfNumbers += res
	}
	fmt.Printf("Cумма цифр в числе %d: %d.\n", naturalNumber, sumOfNumbers)
}