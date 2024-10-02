package main

// 3.5 В бесконечном цикле приложение запрашивает у пользователя числа.
//     Ввод завершается, как только пользователь вводит число "-1".
//     После завершения ввода приложение выводит сумму чисел.
//     Используем конструкцию:
//     for {
//       // body
//     }


import (
    "fmt"
)

func main() {
    var (
        sum int
        number int
    )
    println("Вводите числа")
    outer:
    for {
            fmt.Scan(&number)
            if number == -1 {
                break outer
            }
            sum += number        
        }
    fmt.Printf("Сумма чисел = %v\n", sum)
}
