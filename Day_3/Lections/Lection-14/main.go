// Показать документацию пакета
package main

/*
	1. Явная функция - локально-определенный блок кода, имеющий имя (явное определение)
	Функцию необходимо определить + функцию необходимо вызвать
	2. Сигнатура функий и их определение
	func functionName(params type) typeReturnValue {
		function body
	}
*/


import (
	"fmt"
)

func main()  {
	fmt.Println("call function")
	// 4. Вызов функции
	res := add(10, 20)
	fmt.Println("Result of add(10, 20) is", res)
	fmt.Println("Result of add(10, 20) is", res)

	per, square := rectangleParameters(12, 4)
	fmt.Println("per:", per, "square:", square)

	onlyPerimetr, _ := rectangleParameters(3, 9)
	fmt.Println("onlyPerimetr", onlyPerimetr)

	perNamed, squareNamed := namedReturn(10, 4)
	fmt.Println("per:", perNamed, "square:", squareNamed)

	fmt.Println(funcWithReturn(14, 9))

	emptyReturn(12)

	helloVariadic()
	helloVariadic(70)
	helloVariadic(10, 20, 30, 40)

	someStrings(3, 5)
	someStrings(2, 7, "hello", "age", "golang", "programm")

	resSum1 := sumVariadic(10, 20, 30, 40)
	fmt.Println(resSum1)

	sliceNumbers := []int{5, 7, 8, 9}
	fmt.Println(sumSlice(sliceNumbers))

	resSum2 := sumVariadic(sliceNumbers...)
	fmt.Println(resSum2)

}	

// 3. Простейшая функция (определить функцию можно как до момента ее вызова в функции main),
// так и в любом месте пакета, главное чтобы она была определена в принципе где-нибудь

// add function return sum of two numbers
func add(a int, b int) int {
	result := a + b
	return result
}

// 5. Функция с однотипными параметрами
func mult(a, b, c, d int) int {
	return a * b * c * d // можно указать выражение
}

// 6. Возврат больше чем одного значения (returnType1, returnType2.....)
func rectangleParameters(a, b int) (int32, int) {
	perimetr := int32(2 * (a + b))
	area := a * b
	return perimetr, area
}

// 7. Именованный возврат значений
func namedReturn(width, height int) (perimetr int32, area int) {
	perimetr = int32(2 * (width + height))
	area = width * height
	return // не нужно указывать возвращаемые значения
}

// 8. При вызове оператора return функция прекращает свое выполнение и что-то возвращает
func funcWithReturn(a, b int) (int, bool) {
	if a > b {
		return a - b, true
	}

	if a == b {
		return a, true
	}

	return 0, false
}

// 9. Что вернет функция в случае, если return не указан (или он пустой)
func emptyReturn (a int) {
	fmt.Println("I'm emtpyReturn with parameters...", a)
}

// 10. Variadic parameters (континуальные параметры)
func helloVariadic(a ...int) {
	fmt.Printf("value %v and it's type %T\n", a, a)
}

// 11. Смешивание параметров с variadic (
//   1. Континуальный параметр всегда самый последний
//   2. Variadic parametr - один на всю функцию 
//  )
func someStrings(a, b int, words ...string) {
	fmt.Println("Parametr a:", a)
	fmt.Println("Parametr b:", b)
	fmt.Printf("type words: %T\n", words)

	var count int
	for _, word := range words {
		if len(word) > a && len(word) < b {
			fmt.Println("needed word:", word)
			count++
		}
	}
	fmt.Println("count:", count)
}

// 12. Передача слайса или variadic parameters
func sumVariadic(nums ...int) int {
	var sum int
	for _, number := range nums {
		sum += number
	}
	return sum	
}


func sumSlice(nums []int) int {
	var sum int
	for _, number := range nums {
		sum += number
	}
	return sum	
}