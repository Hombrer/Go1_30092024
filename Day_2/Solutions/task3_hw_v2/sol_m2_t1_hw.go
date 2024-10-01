/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1, 9, 2
Выход: 1, 2, 9

Про sort мы пока ничего не знаем!
Это задача на условный оператор
*/

package main

import "fmt"

func main() {
    var (
        number1 int
        number2 int
        number3 int
    )
    fmt.Print("Введите три числа: ")
    fmt.Scan(&number1, &number2, &number3)
    
    var min, max int

    if number1 < number2 {
        min = number1
        max = number2
    } else {
        min = number2
        max = number1
    }
    
    fmt.Print("В порядке возрастания: ")
    if number3 <= min {
        fmt.Println(number3, min, max)
    } else if min < number3 && number3 <= max {
        fmt.Println(min, number3, max)
    } else {
        fmt.Println(min, max, number3)
    }
}