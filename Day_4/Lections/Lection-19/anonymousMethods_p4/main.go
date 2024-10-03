package main

import (
	"fmt"
)

type University struct {
	city string
	name string
}

// 1. Данный метод связан только со структурой University
func (u *University) FullInfoUniversity() {
	fmt.Printf("Uni name: %s and City: %s\n", u.name, u.city)
}

// 2. В структуру Professor встроены поля структуры University (все методы тоже переходят)
type Professor struct {
	fullName string
	age 	int
	University
}

func (p Professor) Info(){
	fmt.Println("FullName:", p.fullName)
	fmt.Println("Age:", p.age)
}

func main()  {
	prof := Professor {
		fullName: "Boris Bobroff",
		age: 42,
		University: University{
			city: "Moscow",
			name: "BMSTU",
		},
	}

	// 3. Вызываем метод структуры University через экземпляр профессора
	prof.FullInfoUniversity()

	// 4. Вызываем "свой" метод
	prof.Info()
}