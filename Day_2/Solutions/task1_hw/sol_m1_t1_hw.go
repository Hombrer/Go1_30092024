/*
Задача №1
Вход:

    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*Проверка условий по желанию
*/

package main

import "fmt"

const price = 48.00

func main() {
    var (
        distance    int
        consumption int8
    )
    fmt.Print("Введите расстояние: ")
    fmt.Scan(&distance)
    fmt.Print("Введите расход на 100 км: ")
    fmt.Scan(&consumption)

    if (distance >= 50 && 
		distance <= 10000 && 
		consumption >= 5 && 
		consumption <= 25) {
        var cost float64
        cost = (float64(distance) / 100.00) * float64(consumption) * price

        fmt.Println("Стоимость поездки:", cost)
    } else {
        fmt.Println("Неверные входные данные")
    }
}