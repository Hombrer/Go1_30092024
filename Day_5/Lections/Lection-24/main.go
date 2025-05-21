package main

import (
	"fmt"
	"reflect"
)

// САМЫЙ UNSAFE ИНСТРУМЕНТ В GO!!!
type Order struct {
	ordId      int
	customerId int
}

func checkMyType(i interface{}) {
	//Собираюсь в runtime узнать тип того, что было передано в эту функцию
	typeOfSample := reflect.TypeOf(i)
	valueOfSample := reflect.ValueOf(i)
	kindOfSample := typeOfSample.Kind()
	fmt.Println("Type:", typeOfSample)
	fmt.Println("Value:", valueOfSample)
	fmt.Println("Kind:", kindOfSample) // Возвращает глобальное семейство к которому принадлежит экземпляр (структура, интерфейс, функция или примитив)

	// Работает, но медленно
	if reflect.ValueOf(i).Kind() == reflect.Struct {
		v := reflect.ValueOf(i)
		fmt.Println("Number of fields is:", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}

func main() {
	ord := Order{23, 55}
	checkMyType(ord)

	name := "Vasya"
	checkMyType(name)

	checkMyType(func() {})

}
