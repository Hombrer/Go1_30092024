/*
Задача №1.
Программа получает на вход последовательность из 5 целых чисел.

Вам нужно определить вид последовательности:
 - возрастающая
 - убывающая
 - случайная
 - постоянная

В качестве ответа следуют выдать прописными латинскими буквами тип последовательности:
1. ASCENDING (строго возрастающая)
2. WEAKLY ASCENDING (нестрого возрастающая, то есть неубывающая)
3. DESCENDING (строго убывающая)
4. WEAKLY DESCENDING (нестрого убывающая, то есть невозрастающая)
5. CONSTANT (постоянная)
6. RANDOM (случайная)

Примеры входных и выходных данных:
In: 1 2 4 9 11   Out: ASCENDING 2
In: 11 9 4 2 -1  Out: DESCENDING -2
In: 3 8 8 11 12  Out: WEAKLY ASCENDING 1
In: 2 -1 7 21 1  Out: RANDOM not all
In: 5 5 5 5 5    Out: CONSTANT 0

Подсказка: используем метод строки strings.Split()
*/

package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	var asc, desc, constant bool
	var asc_num, desc_num, const_num int
	status := [...]string{
		"", "ASCENDING", "DESCENDING", "RANDOM", "CONSTANT", 
		"WEAKLY ASCENDING", "WEAKLY DESCENDING", "RANDOM",
	}

	fmt.Println("Введите 5 целых чисел:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < 4; i++{
		if arr[i] < arr[i + 1] {
			asc = true
			asc_num = 1
		} else if arr[i] > arr[i + 1] {
			desc = true
			desc_num = 2
		} else {
			constant = true
			const_num = 4
		}
		
	}
	// Первый вариант определения типа последовательности
	index := asc_num + desc_num + const_num
	fmt.Println(status[index])

	fmt.Println("-- Repeat --")

	// Второй вариант определения типа последовательности
	switch  {
	case asc && desc:
		fmt.Println("RANDOM")
	case asc && constant:
		fmt.Println("WEAKLY ASCENDING")
	case asc:
		fmt.Println("ASCENDING")
	case desc && constant:
		fmt.Println("WEAKLY DESCENDING")
	case desc:
		fmt.Println("DESCENDING")
	default:
		fmt.Println("CONSTANT")
	}
}
