package main

import (
	"fmt"
)

// 1. Создадим тип Rectangle
type Rectangle struct {
	length, width int
}

// 2. Создадим конструктор для Rectangle
// Для имен конструктор в Go договорились использовать функции со следующим названием:
// * New() если данный конструктор на файл один (в файле присутствует только одна структура)
// * New<StructName()> - если в файле присутствуют разные структуры

// 3. В Go принято возвращать из конструктора не сам экземпляр, а ссылку на него
func NewRectangle(newLength, newWidth int) *Rectangle {
	return &Rectangle{newLength, newWidth}
}

// 4. Добавим два метода
func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func main()  {
	r := NewRectangle(10, 20)
	fmt.Printf("Type: %T and value %v\n", r, r)
	fmt.Println("Perimeter:", r.Perimeter())
	fmt.Println("Area:", r.Area())
}