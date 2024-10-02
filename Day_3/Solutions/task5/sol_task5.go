package main

// 3.7 Вводим строку без знаков препинания(5 слов).
//     Найти самое длинное слово в строке и вывести сколько в нем букв.

// Пример:
// In: Скажи как дела в учебе
// Out: учебе, 5

// In: Закрепляем циклы в языке Golang
// Out: Закрепляем, 10
// */

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var (
		str string = "Скажи как дела в учебе"
        word string = ""
        max_word string
        max int
        max_word_count int = 1
	)
    str += " " // Для отработки алгоритма с последним словом
    for _ , char := range str {      
        if char != ' ' {   
            word += string(char)
        } else {    
            if utf8.RuneCountInString(word) > max {
                max = utf8.RuneCountInString(word)
                max_word = word
            } else if utf8.RuneCountInString(word) == max {
                max_word_count ++ 
            }
        word = ""
        }            
    }
	fmt.Printf(`
    Самое длинное первое слово = '%v'
    Число символов в этом слове = %v
    Число слов максимальной длинны = %v
    `, max_word, max, max_word_count)
}
