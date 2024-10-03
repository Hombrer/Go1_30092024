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
	"fmt"
	"bufio"
	"os"
	"unicode/utf8"
)

var (
	digits string = "0123456789"
	lowercase string = "abcdefghijklmnopqrstuvwxyz"
	uppercase string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special string = "_!@#$%^&"
)

func main() {
	outer:
	for i := 0; i < 5; i++ {
		fmt.Println("Введите пароль:")
		str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		result := checkPassword(str)
		if result == true {
			fmt.Println("*** Пароль правильный и он принят ***")
			break outer
		} else if i != 4 {
			fmt.Printf("*** Осталось %v попыток ***\n", 4 - i)
		} else {
			fmt.Println("*** Попытки закончились запустите программу снова ***")
			break outer
		}
	}
}

// Проверка пароля на соответствие парольной политки
func checkPassword(pass string) bool {
	status := 0
	if checkPasswordChars(pass, digits) {
		status += 1
	} else {
		fmt.Println("В пароле нет цифр!")
	}

	if checkPasswordChars(pass, lowercase) {
		status += 1
	} else {
		fmt.Println("В пароле нет маленьких букв!")
	}

	if checkPasswordChars(pass, uppercase) {
		status += 1
	} else {
		fmt.Println("В пароле нет больших букв!")
	}

	if checkPasswordChars(pass, special) {
		status += 1
	} else {
		fmt.Println("В пароле нет спец символов!")
	}
	if checkLenPassword(pass) {
		status += 1
	} else {
		fmt.Println("Длина пароля должна быть от 8(мин) до 15(макс) символов")
	}

	return status == 5

}

// Проверка пароля на наличее символла из переданного набора
func checkPasswordChars(password string, chars string) bool {
	flag := false
	outer:
	for _, pass := range password {
		for _, char := range chars {
			if string(char) == string(pass) {
				flag = true
				break outer
			}
		}
	}
	return flag
}

// Проверка пароля на его длинну
func checkLenPassword(pass string) bool {
	flag := false
		if  utf8.RuneCountInString(pass) >= 8 && utf8.RuneCountInString(pass) <= 15 {
			flag = true
		}
	return flag
}