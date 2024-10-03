package main

import (
	"fmt"
)

// 1. Вложенные структуры (вложение структур)
// Это использование одной структуры, как тип поля для в другой стр-ре.
type University struct {
	age       int
	yearBased int
	infoShort string
	infoLong  string
}

type Student struct {
	firstName  string
	lastName   string
	university University
}

// 4. Встроенные структуры (когда мы добавляем поля одно струкуры к другой без имени)
type Professor struct {
	firstName  string
	lastName   string
	age	       int
	greatWorks string
	University // В этом месте происходит добавление всех полей структуры University
}

func main() {
	// 2. Создание экземпляров структур с вложением
	student := Student {
		firstName: "Fedya",
		lastName: "Petrov",
		university: University{
			yearBased: 1991,
			infoShort: "good University",
			infoLong: "It's very good University",
		},
	}
	// 3. Получение доступа к вложенным полям структуры
	fmt.Println("firstName:", student.firstName)
	fmt.Println("lastName:", student.lastName)
	fmt.Println("yearBased of Uni:", student.university.yearBased)
	fmt.Println("infoLong of Uni:", student.university.infoLong)

	// 5. Создание экземпляров с встроенной структурой
	prof := Professor {
		firstName: "Anatoly",
		lastName: "Smirnov",
		age: 35,
		greatWorks: "Ultimate Go programming",
		// papers map[string]string - добавление этого поля делает структуру несравнимой
		University: University{
			yearBased: 1734,
			infoShort: "short Info",
			age: 2024 - 1734,
			infoLong: "long Info",
		},
	}

	// 6. Обращение к состояниям встроенный структуры
	fmt.Println("firstName:", prof.firstName)
	fmt.Println("lastName:", prof.lastName)
	fmt.Println("yearBased of Uni:", prof.yearBased)
	fmt.Println("infoLong of Uni:", prof.infoLong)
	fmt.Println("Age:", prof.University.age) // prof.age - Получим доступ к полю ВЫШЕЛЕЖАЩЕЙ стр-ры

	// 7. Сравнение экземпляров
	profLeft := Professor{}
	profRight := Professor{}
	fmt.Println(profLeft == profRight)

	// 8. Если хотя бы одно из полей структуры - несравнимо - то и вся структура тоже несравнима!
}
