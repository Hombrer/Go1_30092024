package main // он тоже main, но функция main() одна на весь пакет!

// Реализуем 4 функции в моделе calculator

func add(a, b int) int {
	return a + b
}

func Mult(a, b int) int {
	return a * b
}

func Sub(a, b int) int {
	return a - b
}

func Div(a, b int) int {
	return a / b
}