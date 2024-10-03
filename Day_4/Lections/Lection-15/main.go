package main

import (
	"fmt"
)

func main()  {
	// 1. Анонимная функция
	anon := func(a, b int) int {
		return a + b
	}

	fmt.Println("Anon:", anon(10, 20))
	fmt.Println("biFunction(30, 40):", bigFunction(30, 40))

	var command string = "substraction"
	res := calcAndReturnValidFunc(command)
	fmt.Println("Result with command", command, "is", res(10, 20))

	// 4. Тип функции
	fmt.Printf("Тип функции res: %T\n", res)
	fmt.Printf("Тип функции calcAndReturnValidFunc: %T\n", calcAndReturnValidFunc)

	fmt.Println("Call receiveFuncAndReturnValue(add):", receiveFuncAndReturnValue(add))
	fmt.Println("Call receiveFuncAndReturnValue(add):", receiveFuncAndReturnValue(add))
	fmt.Println(
		"Call receiveFuncAndReturnValue(random):", 
		receiveFuncAndReturnValue(func(a, b int) int {return a*a + b*b + 2*a*b}))
}

// 2. Анонимная функция внутри явной
func bigFunction(aArg, bArg int) int {
	return func(a, b int) int {return a + b + 100}(aArg, bArg)
}

// 3. Явные функции в Go (в принципе любая функция) - это экземпляр 1-го уровня,
// поэтому ее можно присваивать переменной, передавать в качестве параметра 
// и возвращать в качестве результата
// Пример возврата функции
func calcAndReturnValidFunc(command string) func(a, b int) int {
	if command == "addition" {
		return func(a, b int) int {return a + b}
	} else if command == "substraction" {
		return func(a, b int) int {return a - b}
	} else {
		return func(a, b int) int {return a * b}
	}
}

// 5. Функция как входящий параметр в другой функции
func receiveFuncAndReturnValue(f func(a, b int) int) int {
	var intVarA, intVarB int
	intVarA = 100
	intVarB = 200

	return f(intVarA, intVarB)
}

func add(a, b int) int {
	return a + b
}