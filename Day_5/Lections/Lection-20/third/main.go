package main

import (
	"fmt"
)

type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}


func (st Student) Describe() {
	fmt.Printf("%s and %d years old\n", st.name, st.age)
}

func (st Student) GetName() {
	fmt.Println("Name:", st.name)
}

func TypeFinder(i interface{}) {
	// type assertion together with switch
	// Присваиваем переменной v значение лежащие под предполагаемым интерфейсом
	switch v := i.(type) {
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("This is int")
	case Describer: // Вывод - с типом интерфейса можно
		fmt.Println("I'm is Describer")
		v.Describe()
	default:
		fmt.Println("Unknown type")
	}
}

func FindStudent(i interface{}) {
	// type assertion together with switch
	switch v := i.(type) {// Присваиваем переменной v значение лежащие под предполагаемым интерфейсом
	case Student:
		fmt.Println("I'm student")
		v.GetName() // Используем метод, который есть только у Student
	default:
		fmt.Println("Unknown type")
	}
}


func main()  {
	std := Student{"Alex", 23}
	TypeFinder(10)
	TypeFinder("Hello")
	TypeFinder(std)
	TypeFinder(nil)
	var desc Describer
	desc = Student{"Mark", 48}
	FindStudent(desc)
	FindStudent("Hello")
}