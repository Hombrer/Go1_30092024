// Задача №1
// Программа запрашивает имя пользователя и возраст
// Нужно вывести имя и возраст за вычетом одного года
package main

import "fmt"

const minus = 1

func main() {
	var (
		age int // 0
		name string // ""
	)
	fmt.Println("Enter name:")
	fmt.Scan(&name)
	fmt.Println("Enter age:")
	fmt.Scan(&age)
	fmt.Printf("Your name is %s.\nYour age is %d.\n", name, age - minus)
}