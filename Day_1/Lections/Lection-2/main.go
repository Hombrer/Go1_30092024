/*
	Пример мнострочного комментария.
	Примеры из Lection-2
*/
package main

import "fmt"


func main() {
	// Пример однострочного комментария. Ctrl + /
	// var <variable_name> <variable_type> + auto initialization
	var age int
	fmt.Println("My age is:", age)

	var height int = 183 // + manual initialization
	fmt.Println("My height is:", height)

	// В чем "полустрогость" типизации?
	var uid = 12345
	fmt.Println("My uid:", uid)
	fmt.Printf("%T\n", uid)

	// Константы
	const price, tax float32 = 275.00, 27.50 // Типизированные константы
	const quantity, inStock = 2, true // Нетипизированные константы
	fmt.Println("Total:", 2 * quantity * (price + tax))

	// Variables
	var cost float32 = 275.00
	var total float32 = 27.50
	cost = 300
	fmt.Println(cost + total)

	// Error block
	// var value = 275.00
	// var taxes float32 = 27.50
	// fmt.Println(value + taxes)

	// Work block
	var value = 275.00
	var taxes float32 = 27.50
	fmt.Println(value + float64(taxes))
}