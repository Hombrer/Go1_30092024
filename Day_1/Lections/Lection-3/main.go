package main

import (
	"fmt"
	"math"
)

func main(){
	var result float64
	var d, e float64 = 10, 4
	// Результат целое частное, потому что 2 целых числа
	var check float64 = 20 / 3 

	a := 42
	b := 20
	c := a / b
	result = d / e
	resultNew := d / 2

	fmt.Printf("type of c %T and value of c %v\n", c, c)
	fmt.Printf("type of result %T and value of result %v\n", result, result)
	fmt.Printf("type of resultNew %T and value of resultNew %v\n", resultNew, resultNew)
	fmt.Printf("type of check %T and value of check %v\n", check, check)

	fmt.Print("Инкремент: ")
	// check = check + 1
	check++
	fmt.Println("check =", check)
	fmt.Print("Декремент: ")
	// c = c - 1
	c--
	fmt.Println("c =", c)

	// Операция взятия остатка
	newCheck := int(check) % 3
	fmt.Println("newCheck:", newCheck)

	// Операция возведения в степень из модуля math
	total := math.Pow(2, 10)
	fmt.Printf("Total = %.1f\n", total)
}

