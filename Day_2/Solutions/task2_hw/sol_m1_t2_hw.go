/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример: 
вход: 346, выход: 643
вход: 120, выход: 021
вход: 100, выход: 001
*/

package main

import (
	"fmt"
)

func main() {
	var number int32

	fmt.Println("Введите 3х значное число")
	fmt.Scan(&number)
	first := number/100
	second := number%100/10
	last := number%10
	new_number := last*100+second*10+first
	fmt.Printf("Наше число = %03d ,  реверсная запись трехзначного числа = %03d \n",number, new_number)
}