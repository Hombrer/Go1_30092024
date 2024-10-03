/*
Задача №2

Вход:
Пользователь должен ввести "правильный пароль", состоящий из:
цифр, букв латинского алфавита(строчные и прописные) и
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "_!@#$%^&"

Выход:
Написать, что пароль правильный и он принят.

Пример:
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8

Реализацию оформить через функцию:
1. checkPassword(pass string) (bool, errors <- на усмотрение)
*/
package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// special = "_!@#$%^&"
var specialCharacters = map[rune]int{
	'_': 0,
	'!': 0,
	'@': 0,
	'#': 0,
	'$': 0,
	'%': 0,
	'^': 0,
	'&': 0,
}

const tryCount int = 5

func main() {
	var count int = 0

	for count < tryCount {
		fmt.Print("Введите пароль: ")
		var pass string
		fmt.Scan(&pass)

		ok, error := checkPassword(pass)

		if ok {
			fmt.Println("Пароль принят")
			break
		} else {
			fmt.Println(error.Error())
		}

		count++
	}
}

func checkPassword(pass string) (bool, error) {
	passLength := utf8.RuneCountInString(pass)

	if passLength < 8 {
		return false, errors.New("Длина пароля должна быть меньше 8 символов")
	}

	if passLength > 15 {
		return false, errors.New("Длина пароля должна не должна быть больше 15 символов")
	}

	var flags byte

	for _, symbol := range pass {
		if isDigit(symbol) {
			flags |= 1
		} else if isLowerCaseLetter(symbol) {
			flags |= 2
		} else if isUpperCaseLetter(symbol) {
			flags |= 4
		} else if isSpecialCharacter(symbol) {
			flags |= 8
		}
	}

	if flags&1 == 0 {
		return false, errors.New("В пароле отсутствует цифра")
	}

	if flags&2 == 0 {
		return false, errors.New("В пароле отсутствует символ в нижнем регистре")
	}

	if flags&4 == 0 {
		return false, errors.New("В пароле отсутствует символ в верхнем регистре")
	}

	if flags&8 == 0 {
		return false, errors.New("В пароле отсутствует спец. символ")
	}

	return true, nil
}

func isDigit(input rune) bool {
	var code = int(input)

	return code >= 48 && code <= 57
}

func isLowerCaseLetter(input rune) bool {
	var code = int(input)

	return code >= 97 && code <= 122
}

func isUpperCaseLetter(input rune) bool {
	var code = int(input)

	return code >= 65 && code <= 90
}

func isSpecialCharacter(input rune) bool {
	_, ok := specialCharacters[input]

	return ok
}