package main

import "fmt"

// 1. Данная функция будет запущена в качестве горутины.
// Важно: горутины никогда ничего не возвращают через явное использование return
func newGoRoutine() {
	fmt.Println("Hey, I'm new Gorutine!")
}

// 2. функция main - на самом деле тоже горутина.
// Особенность в том - что если эта горутина завершается - все остальные запущенные убиваются сразу!
func main() {
	go newGoRoutine() // в этот момент происходит формирование запроса на вызов функции в отдельной горутины.
	// соответственно код основной горутины main продолжает сразу же выполняться и не ждет завершения newGoRoutine()
	fmt.Println("Main goroutine work!")
	//Запустим код и....
}
