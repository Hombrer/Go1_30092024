package main

import (
	"fmt"
	"unicode/utf8"
)

// 1. Описание интерфейса
type BigWorld interface {
	IsBig() bool
}

// 2. Наш тип-претендент, который должен иметь метод IsBig()
type MySuperString string

// 3. Реализация IsBig()
func (mss MySuperString) IsBig() bool {
	if utf8.RuneCountInString(string(mss)) > 10 {
		return true
	}
	return false
}

func main() {
	sample := MySuperString("hello")
	var interfaceSample BigWorld // объявление переменной с типом BigWorld
	interfaceSample = sample // так можно сделать, т.к. sample отвечает интерфейсу BigWorld
	fmt.Println("IsBig?", interfaceSample.IsBig())

	// 4. Попытка присвоить значение с типажом, не удовлетворяющему интерфейсу приводит к ошибке
	// var interfaceBadSample BigWorld
	// interfaceBadSample = "long string" // string не имеет реализации IsBig()

}