package main

import (
	"fmt"
)

// 1. А что если создать интерфейс, в котором вообще нет никаких требований к поведению
type Empty interface{}

// 2. Вопрос в том, кто удовлетворяет этому интерфейсу. Ответ - кто угодно

// 3. Реализуем функцию, которая может принимать что угодно
func Describer(pretendent interface{}){
	fmt.Printf("Type = %T and value %v\n", pretendent, pretendent)
}

type Student struct {
	name string
}

// 4. Type Assertion
func SemiGeneric(pretendent interface{}) {
	value, ok := pretendent.(int) // Проверка того, что pretendent типа int
	fmt.Println("Value:", value, "OK?", ok)
}

func main()  {
	strSample := "Hello world"
	// 5. Передача параметра без предварительного присваивания в промежуточную переменную
	Describer(strSample)

	intSample := 200
	Describer(intSample)

	Describer(Student{"Vasya"})

	// 6. Работа с полу-дженериком
	SemiGeneric(10)
	SemiGeneric(Student{"Fedya"})
	SemiGeneric("Hello students")

}