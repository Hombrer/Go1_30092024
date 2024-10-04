package main

import (
	"fmt"
)

/*
	0. Структура - явно декларированный заименованный набор СОСТОЯНИЙ.
	Структура, исходя из своего названия, отвечает на вопрос - из ЧЕГО я должна состоять,
	чтобы считаться тем ТИПОМ, который декларируется структурой.
	Структура описывает ПАТТЕРН СОСТОЯНИЯ.

	1. Интерфейс - явно декларированный набор сигнатур ПОВЕДЕНИЙ (чаще всего в виде набора методов),
	удовлетворив который, может считаться типом, который декларирует интерфейс.
	Отвечает на вопрос - что я должен уметь делать, чтобы считаться тем типом,  который декларирует интерфейс.
	Интерфейс описывает ПАТТЕРН ПОВЕДЕНИЯ
*/

// 2. Объявление интерфейса
type FigureBuilder interface {
	// Набор сигнатур методов, которые необходимо реализовать в структуре-претенденте
	// Во-первый, у претендента должен быть метод Area() возвращающий int
	Area() int
	// Во-вторых, должен быть метод Perimeter() возвращающий int
	Perimeter() int
}


// 3. Декларируем структуры-претенденты
// 3.1 Первый - прямоугольник
// Когда методы Area() и Perimeter() реализованы, говорят, что Rectangle удовлетворяет
// условиям интерфейса FigureBuilder
// Rectangle реализует интерфейс FigureBuilder
type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

// 3.2 Первый - окружность
// Когда методы Area() и Perimeter() реализованы, говорят, что Circle удовлетворяет
// условиям интерфейса FigureBuilder
// Circle реализует интерфейс FigureBuilder
type Circle struct {
	radius int
}

func (c Circle) Area() int {
	return 3 * c.radius * c.radius
}

func (c Circle) Perimeter() int {
	return 2 * 3 * c.radius
}


func main()  {
	// 4. Создадим по 3 экземпляра
	r1 := Rectangle{10, 20}
	r2 := Rectangle{30, 50}
	r3 := Rectangle{1, 1}
	c1 := Circle{5}
	c2 := Circle{10}
	c3 := Circle{15}

	// 5. Посчитаем общую площадь этих фигур
	rectangles := []Rectangle{r1, r2, r3}
	totalSumOfRectangles := 0
	for _, rect := range rectangles {
		totalSumOfRectangles += rect.Area()
	}

	circles := []Circle{c1, c2, c3}
	totalSumOfCircles := 0
	for _, circle := range circles {
		totalSumOfCircles += circle.Area()
	}

	fmt.Println("Total area is", totalSumOfRectangles + totalSumOfCircles)

	// 6. Теперь воспользуемся интерфейсом, который мы создали ранее
	// Объявляю слайс экземпляров, удовлетворяющих интерфейсу FigureBuilder
	figures := []FigureBuilder{r1, r2, r3, c1, c2, c3}

	// 7. Посчитаем общую площадь
	total := 0
	for _, figure := range figures {
		total += figure.Area()
	}
	fmt.Println("Repeat total area is", total)

	// 8. У каждого элемента из слайса FigureBuilder можно 100% вызвать метод Area(), 
	// который вернет int или Perimeter(), который тоже вернет int.

	// 9. Говорят, что интерфейсы - это полу-абстрактный тип данных.


}