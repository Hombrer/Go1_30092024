/* 
	Задача №1
	Написать функцию, которая расшифрует строку. 
	code = "220411112603141304"
	Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
	Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
	Вход: строка из цифр. Выход: Текст.
	Проверка работы функции выполняется через другую строку.

	Задача №1.1
	Реализовать и функцию зашифровки

	codeToString(code) -> "???????' 

	stringToCode("hello") -> "??????"
*/

package main

import (
	"fmt"
	"bufio"
	"os"
)
func main() {
	var (

		abc string = "abcdefghijklmnopqrstuvwxyz "
	)
	mapCode := make(map[string]string)
	mapDigits := make(map[string]string)
	// Создаем map c кодами
	for i, char := range abc {
		mapCode[fmt.Sprintf("%02d", i)] = fmt.Sprintf("%v", string(char))
		mapDigits[fmt.Sprintf("%v", string(char))] = fmt.Sprintf("%02d", i)
	}
	fmt.Println("Наша мапа для расшифровывания", mapCode)
	fmt.Println("Наша мапа для зашифровывания", mapDigits)

	fmt.Println("Введите строку для шифровки:")
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Ваша зашифрованная строка:", stringToCode(str, mapDigits))

	fmt.Println("Введите строку для расшифровки:")
	str2, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Ваша расшифрованная строка:", codeToString(str2, mapCode))
}

func stringToCode(str string, mapCode map[string]string) string {
	var answer string
	for _, char := range str {
		answer += mapCode[string(char)]
	}
	return answer
}

func codeToString(str string, mapCode map[string]string) string {
	var answer string
	for i := range str {
		if i % 2 != 0 {
			key_symbol := string(str[i-1]) + string(str[i])
			answer += mapCode[key_symbol]
		}
	}
	return answer
}