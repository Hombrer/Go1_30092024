package main

import (
	"fmt"
)

type Employee struct {
	name 	 string
	position string
	salary 	 int
	currency string
}

// 1. Методы - функции, привязанные к определенным структурам
// Шаблон
// func (s Struct) MethodName(parameters type) output type(s) { body }
//     Receiver - получатель метода
func (e Employee) DisplayInfo() {
	fmt.Println("Name:", e.name)
	fmt.Println("Position:", e.position)
	fmt.Printf("Salary: %d %s\n", e.salary, e.currency)
}


func main()  {
	// 2. Создание экземпляра
	emp := Employee {
		name: "Mark",
		position: "Senior Golang programmer",
		salary: 200_000,
		currency: "RUB",
	}

	// 3. Вызов метода
	emp.DisplayInfo()
}