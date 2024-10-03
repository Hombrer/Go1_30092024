package main

import (
	"fmt"
)

// 1. Структура - заименованный набор полей(состояний), определяющий новый тип данных


// 2. Определение структуры(явное определение)
type Student struct {
	firstName   string
	lastName    string
	age			int
}


// 11. Структура с анонимными полями
type Human struct {
	firstName   string
	lastName    string
	string
	int
	bool
}


// 3. Если имеется ряд состояний одинакового типа, то можно сделать так
type AnotherStudent struct {
	firstName, lastName, groupName  string
	age, courseNumber 				int
}

func PrintStudent(std Student) {
	fmt.Println("=======================")
	fmt.Println("FirstName:", std.firstName)
	fmt.Println("LastName:", std.lastName)
	fmt.Println("Age:", std.age)
}


func main()  {
	// 4. Создание представителя(экземпляра) структуры
	student := Student{
		firstName: "Fedya",
		age: 	   21,	
		lastName:  "Petrov",
	}
	PrintStudent(student)

	student2 := Student{"Petr", "Ivanov", 19} // порядок указания свойств - такой же как и в стр-ре
	PrintStudent(student2)

	// 5. Что если не все поля структуры определены?
	student3 := Student{
		firstName: "Vasya",
	}
	PrintStudent(student3)

	// 6. Анонимные структуры (структуры без имени)
	anonStudent := struct {
		age     		int
		groupID 		int
		professorName 	string
	}{
		age:	23,
		groupID: 2,
		professorName: "Alekseev",

	}
	fmt.Println("AnonStudent:", anonStudent)

	// 7. Доступ к полям(состояниям) и их модификация
	studMark := Student{"Mark", "Ivanov", 19}
	fmt.Println("FirstName:", studMark.firstName)
	fmt.Println("LastName:", studMark.lastName)
	fmt.Println("Age:", studMark.age)
	studMark.age += 2
	fmt.Println("New age:", studMark.age)

	// 8. Инициализация пустой структуры
	emptyStudent1 := Student{}
	var emptyStudent2 Student
	PrintStudent(emptyStudent1)
	PrintStudent(emptyStudent2)

	// 9. Указатели на экземпляры структур
	studPointer := &Student{
		firstName: "Igor",
		lastName: "Sidorov",
		age: 23,
	}
	fmt.Println("Value of studPointer:", studPointer)

	secondPointer := studPointer
	(*secondPointer).age += 20
	fmt.Println("Value of studPointer after changing:", studPointer)
	studPointerNew := new(Student)
	fmt.Println(studPointerNew)

	// 10. Работа с доступом к полям через указатель
	fmt.Println("Age via (*...).age:", (*studPointer).age)
	// Неявно происходит разыменование и обращение к соответствующему полю
	fmt.Println("Age via .age:", studPointer.age)

	// 12. Создание экземпляра с анонимными полями
	human := &Human {
		firstName: "Bob",
		lastName: "Smirnoff",
		string: "Additional info",
		int:   -1,
		bool:  true,
	}
	fmt.Println(human)
	fmt.Println("Anon field string:", human.string)
}

