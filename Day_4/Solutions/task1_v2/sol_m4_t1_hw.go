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
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func codeToString(code string) (string, error) {
	var result strings.Builder

	// Проходим по строке, беря каждые две цифры
	for i := 0; i < len(code); i += 2 {
		num, err := strconv.Atoi(code[i : i+2]) // Преобразуем подстроку в число
		if err != nil {
			return "", err
		}

		// Преобразуем число в символ
		if num >= 0 && num <= 25 {
			result.WriteByte(byte(num + 'a')) // 'a' имеет код 97
		} else if num == 26 {
			result.WriteByte(' ') // Пробел
		} else {
			return "", fmt.Errorf("недопустимый код: %d", num)
		}
	}

	return result.String(), nil
}

func stringToCode(text string) (string, error) {
	var result strings.Builder

	for _, char := range text {
		if char == ' ' {
			result.WriteString("26") // Пробел кодируется как 26
		} else if 'a' <= char && char <= 'z' {
			result.WriteString(fmt.Sprintf("%02d", char-'a')) // Кодируем букву
		} else {
			return "", fmt.Errorf("недопустимый символ: %c", char)
		}
	}

	return result.String(), nil
}

func main() {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите код для расшифровки: ")
	if reader.Scan() {
		code := reader.Text()
		decodedText, err := codeToString(code)
		if err != nil {
			fmt.Println("Ошибка расшифровки:", err)
			return
		}
		fmt.Printf("Расшифрованный текст: '%s'\n", decodedText)
	}

	fmt.Print("Введите текст для шифрования: ")
	if reader.Scan() {
		text := reader.Text()
		encodedCode, err := stringToCode(text)
		if err != nil {
			fmt.Println("Ошибка шифрования:", err)
			return
		}
		fmt.Printf("Зашифрованный код: '%s'\n", encodedCode)
	}
}