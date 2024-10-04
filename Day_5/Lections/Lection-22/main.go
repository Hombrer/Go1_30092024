package main // говорим, что в пакете определена функция main()

import (
	"Go1/Day_5/Lections/Lection-22/rectangle"
	// "Lection-22/rectangle" // Прописываем путь до пакета rectangle
	"fmt"
)

// 0. Разделяющая область видимости уровня одного пакета
// 1. В этом примере 2 модуля main и calculator помещены в одну директорию(находятся в одном пакете)

// 2. Разделяющая область видимости
// Все, что в принципе находится внутри одного пакета - доступно из любого модуля без импортирования

func main()  {
	// Данные функции видны компилятору, потому что они в рамках пакета main
	resAdd := add(10, 20) // здесь импорты не нужны!
	fmt.Println("resAdd:", resAdd) 

	resSub := Sub(50, 12)
	fmt.Println("resSub:", resSub) 

	resMult := Mult(3, 9)
	fmt.Println("resMult:", resMult) 

	resDiv := Div(29, 4)
	fmt.Println("resDiv:", resDiv) 

	// Если имя сущности написано с маленькой буквы(переменная, функция, структура, поле структуры, ...) 
	// то ее нельзя передать за пределы пакета.

	// Если сущность написана с большой буквы - ее можно экспортировать за пределы пакета

	// Как запустить
	// go run main.go calculator.go
	// go build main.go calculator.go
	// go install main.go calculator.go
	// go run . // в текущей директории

	// Создадим дополнительный паке с прямоугольником
	// mkdir rectangle && cd rectangle && touch rectangle.go
	r := rectangle.New(10, 20, "green")
	fmt.Println("Perimeter", r.Perimeter())
	
	// Если go.mod уже создан, то облегчается сборка
	// go build (компилятор находит go.mod и видит отправную точку)
	// go install

	// Как правильно, при инициализации проекта, используют go.mod


}