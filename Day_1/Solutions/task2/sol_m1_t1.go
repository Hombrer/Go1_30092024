// Задача №2
// Вход: трехзначное число. 
// Найти первую, вторую и последнюю цифры заданного трехзначного числа.
package main


import "fmt"

func main() {

	var number uint16
	fmt.Println("Введите трехзначное число")
	fmt.Scan(&number)

	fmt.Println("Первая цифра:", number / 100)
	fmt.Println("Вторая цифра:", number / 10 % 10)
	fmt.Println("Третья цифра:", number % 10)
}