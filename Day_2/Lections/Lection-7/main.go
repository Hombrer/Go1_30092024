package main

import (
	"fmt"
	"strings"
)



func main()  {
	// Шаблон цикла for как for
	// for init; condition; post {
	// init - блок инициализации переменных цикла
	// condition - условие (если условие верно - тело цикла выполняется,
	// если - нет, цикл завершает свою работу)
	// post - изменение переменной цикла (инкрементарное или декрементарное действие)
	// }

	for i := 0; i <= 5; i++ {
		fmt.Println("Current value of i:", i)
	}

	// Важный момент. В качестве init можно использовать выражение ТОЛЬКО через :=

	// break - команда, которая прерывает текущее выполнение цикла и 
	// передает управление следующим за циклом инструкциям
	for i := 11; i <= 22; i++ {
		if i > 20 {
			break
		}
		fmt.Println("Current value of i:", i)
	}
	fmt.Println("After for loop with BREAK")

	// continue - команда, которая прерывает текущее выполнение цикла и 
	// передает управление следующей итерации цикла
	for i := 105; i <= 120; i++ {
		if i > 110 && i <= 115 {
			continue
		}
		fmt.Println("Current value of i:", i)
	}
	fmt.Println("After for loop with CONTINUE")

	// Вложенные циклы и лейблы
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("Выше должен быть треугольник")

	// Бывают ситуции, когда нужно прервать сразу оба цикла
	// Здесь помогут лейблы. Лейблы - это синтаксический сахар

outer:
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j:%d\n", i, j, i+j)
			if i == j {
				break outer // Хочу, чтобы все циклы остановились, даже внешний
			}
		}
	}
	fmt.Println("After outer stop")

	// Шаблон цикла for как while (модификация for)
	// 1. Классический цикл while do
	var loopVar int = 0
	// while (loopVar < 10) {
	// 	  ....
	// loopVar++
	// }

	for loopVar < 10 {
		fmt.Println("In while like loop", loopVar)
		loopVar++
	}
	// loopVar - это переменная цикла, а не main
	// Этот пример: можно но не нужно
	// for loopVar := 0;loopVar < 10; {
	// 	fmt.Println("In while like loop", loopVar)
	// 	loopVar++
	// }
	// fmt.Println(loopVar) // 0

	// 2. Классический бесконечный цикл
	var password string
outher2:
	for {
		fmt.Println("Insert password:")
		fmt.Scan(&password)
		if strings.Contains(password, "1234"){
			fmt.Println("Weak password. Try again.")
		} else {
			fmt.Println("Password accepted")
			break outher2
		}
	}

	// 3. Цикл с множественными переменными цикла
	for x, y := 0, 1; x < 5 && y < 8; x, y = x+1, y+2 {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}

	// 4. x1 - это переменная main, а не цикла
	x1 := 5
	for x1 = 10; x1 < 15; x1 += 1 {
		fmt.Printf("%d\n", x1)
	}
	fmt.Println("Result:", x1)
}