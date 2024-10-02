package main

import (
	"fmt"
	"strings"
)

func main()  {
	// 1. Map - это набор ключ:значение
	// Инициализация пустой мапы
	mapper := make(map[string]int)
	fmt.Println("Empty map:", mapper)

	// 2. Добавление элементов в мапу
	mapper["Petr"] = 23
	mapper["Elena"] = 12
	fmt.Println("After adding pairs in map:", mapper)

	// 3. Инициализация мапы с указанием пар
	newMapper := map[string]int{
		"Alice": 192,
		"Bob": 134,
	}	
	newMapper["Nikolay"] = 87
	fmt.Println("NewMapper:", newMapper)

	// 4. Что может быть ключом в мапе?
	// 4.1 ВАЖНО: Мапа в Golang не упорядочена
	// 4.2 Ключом в мапе может быть только сравнимый тип (==, !=)

	// 5. Нулевое значение для map
	// var mapZeroValue map[string]int // mapZeroValue == nil
	// mapZeroValue["Alice"] = 12

	// 6. Получение значения из мапы
	testPerson := "Alice"
	// 6.1 Получение элементы, который есть в мапе
	fmt.Println("Salary of Alice:", newMapper[testPerson])
	// 6.2 Получение элементы, которого нет в мапе
	testPerson = "Vasya"
	fmt.Println("Salary of Vasya:", newMapper[testPerson])

	// 7. Проверка нахождения ключа
	employee := map[string]int{
		"Denis": 0,
		"Alice": 0,
		"Vasya": 0,
	}	
	
	// 7.1 При обращении по ключу - возвращаются 2 значения
	if value, ok := employee["Denis"]; ok {  // ;ok -> ;ok == True
		fmt.Println("Denis exists and it's value:", value)
	} else {
		fmt.Println("Denis does not exits in map.")
	}

	if value, ok := employee["Bob"]; ok {
		fmt.Println("Bob exists and it's value:", value)
	} else {
		fmt.Println("Bob does not exits in map.")
	}

	// 8. Перебор значений мапы
	fmt.Println(strings.Repeat("=", 15))
	for key, value := range employee {
		fmt.Printf("key: %s and value: %v\n", key, value)
	}

	// 9. Как удалять пары
	// 9.1 Удаление существующей пары
	fmt.Println("Before deleting", employee)
	delete(employee, "Vasya")
	fmt.Println("After deleting", employee)

	// 9.2 Удаление несуществующей пары
	if _, ok := employee["Anna"]; ok {
		delete(employee, "Anna") // Очень дорогая операция
	}

	// 10. Количество пар == длина мапы
	fmt.Println("Pairs amount in map:", len(employee))

	// 11. Мапа(как и слайс) - это ссылочный тип
	words := map[string]string{
		"One": "Один",
		"Two": "Два",
	}
	newWords := words
	newWords["Three"] = "Три"
	delete(newWords, "One")
	fmt.Printf("words: %v and it's pointer: %p\n", words, &words)
	fmt.Printf("newWords: %v and it's pointer: %p\n", newWords, &newWords)


	// 12. Сравнение мап
	// 12.1 Сравнение массивов (массив можно использовать как ключ в мапе)
	if [3]int{1, 2, 3} == [3]int{1, 2, 3} {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")

	}
	// 12.2 Сравнение слайсов. 
	// Из-за того, что слайсы ссылочный тип - его можно сравнивать на равенство только с nil.
	// if []int{1, 2, 3} == []int{1, 2, 3} {
	// 	fmt.Println()
	// }

	// 12.3 Сравнение мап
	// Из-за того, что мапы ссылочный тип - его можно сравнивать на равенство только с nil.
	aMap := map[string]int{
		"a": 1,
	}
	var bMap map[string]int

	if aMap == nil {
		fmt.Println("Zero value map")
	}

	if bMap == nil {
		fmt.Println("Zero value map of bMap")
	}

	// Следствие. 
	// Если слайс или мапа являются составляющими какой-либо структуры - 
	// структура автоматически не сравнивма
}