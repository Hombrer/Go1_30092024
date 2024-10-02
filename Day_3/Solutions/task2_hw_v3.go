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
)

func main() {
    var digits [5]int
    
	// Побитовые операции
	// & - bit and
	// | - bit or
	// ^ - bit xor
	// >> - сдвиг вправо
	// << - сдвиг влево

    fmt.Print("Введите последовательность из 5 чисел через пробел: ")

    fmt.Scan(&digits[0], &digits[1], &digits[2], &digits[3], &digits[4])

    var flags byte

    for i := 0; i < len(digits) - 1; i++ {
        digit1 := digits[i]
        digit2 := digits[i + 1]
        
        if digit1 < digit2 {
            flags |= 1 // flags = flags | 1, XX0 | 001 = XX1
        } else if digit1 > digit2 {
            flags |= 2 // X0X | 010 = X1X
        } else {
            flags |= 4 // 0XX | 100 = 1XX
        }
    }
    
    var result string
    
    switch flags {
    case 1:
        result = "ASCENDING"
    case 2:
        result = "DESCENDING"
    case 4:
        result = "CONSTANT"
    case 5:
        result = "WEAKLY ASCENDING"
    case 6:
        result = "WEAKLY DESCENDING"
    default:
        result = "RANDOM"
    }
    
    fmt.Print(result)
}