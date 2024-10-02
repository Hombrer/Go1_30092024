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
7. RANDOM (случайная)

Примеры входных и выходных данных:
In: 11 9 4 2 -1  Out: DESCENDING
In: 3 8 8 11 12  Out: WEAKLY ASCENDING
In: 2 -1 7 21 1  Out: RANDOM
In: 5 5 5 5 5     Out: CONSTANT

Подсказка: используем метод строки strings.Split()
*/

package main

import (
    "fmt"
	"bufio"
    "os"
	"strings"
	"strconv"
)
var (
	str string
	number [5]int64
)

func main() {
	println("Введите последовательность из 5 целых чисел:")
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	words := strings.Fields(str) // На выходе массив из строк	
	for i , word := range words{
		number[i], _ = strconv.ParseInt(string(word), 10, 64)
	}

	switch {
	case number[0] == number[1] &&
		 number[1] == number[2] &&
		 number[2] == number[3] &&
		 number[3] == number[4]:
		fmt.Println("CONSTANT")
	case number[0] < number[1] &&
		 number[1] < number[2] &&
		 number[2] < number[3] &&		 
		 number[3] < number[4]:
		fmt.Println("ASCENDING")
	case (number[0] < number[1] || number[0] == number[1]) &&
		 (number[1] < number[2] || number[1] == number[2]) &&
		 (number[2] < number[3] || number[2] == number[3]) &&
		 (number[3] < number[4] || number[3] == number[4]):
	    fmt.Println("WEAKLY ASCENDING")		
	case number[0] > number[1] &&
		 number[1] > number[2] &&
		 number[2] > number[3] &&
		 number[3] > number[4]:
		fmt.Println("DESCENDING")
	case (number[0] > number[1] || number[0] == number[1]) &&
		 (number[1] > number[2] || number[1] == number[2]) &&
		 (number[2] > number[3] || number[2] == number[3]) &&
		 (number[3] > number[4] || number[3] == number[4]):
		fmt.Println("WEAKLY DESCENDING")
	default:
		fmt.Println("RANDOM")
	}
}