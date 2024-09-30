package main

import (
	"fmt"
	"strings"
)

func main() {
	// Boolean => default: false
	var firstBoolean bool
	fmt.Println(firstBoolean)

	// Boolean operands
	aBoolean, bBoolean := true, true
	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT:", !aBoolean)
	
	// Классический условный оператор
	/*
	if condition {
		// body
		}
	*/

	// Классический условный оператор с блоком else
	/*
	if condition {
		// body 1
	} else {
		// body 2 
	}
	*/

	var value int
	fmt.Scan(&value)

	if value % 2 == 0 {
		fmt.Println("Number", value, "is even.")
	} else {
		fmt.Println("Number", value, "is odd.")
	}
	
	// Классический условный оператор с блоками else if
	/*
	if condition1 {
		// body 1
	} else if condition2 {
		// body 2 
	} else if condition3 {
		// body 3
	} else {
		// body 4
	}
	*/

	var color string
	fmt.Scan(&color)

	if strings.Compare(color, "green") == 0 {
		fmt.Println("Color is green")
	} else if strings.Compare(color, "red") == 0 {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Unknown color")
	}

	// Good Инициализация в блоке условного оператора
	// Блок присваивания - только :=
	// Инициализируемая переменная видно ТОЛЬКО внутри области
	// видимости условного оператора (в телах if, else if или else),
	// но не за его пределами.
	if num := 10; num % 2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}

	// Не идеоматично
	if width := 100; width > 100 {
		fmt.Println("width > 100")
	} else {
		fmt.Println("width <= 100")
	}
	// Странное правило номер 1: в Go стараются избегать блоков ELSE
	// Идеоматично
	if width := 100; width > 100 {
		fmt.Println("width > 100")
		return
	} 
	fmt.Println("width <= 100")

}