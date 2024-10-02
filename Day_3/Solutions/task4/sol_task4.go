package main

// 3.6 Получить от пользователя натуральное число, посчитать сумму цифр в нём. 
//     Решить с помощью индексов, т.е. работаем с числом, как со строкой.


import (
    "fmt"
    "strconv"
)

func main() {
	var (
		number string
        sum int64
        n int64
	)
	fmt.Println("Введите натуральное число число")
	fmt.Scan(&number)
    for _, char := range number {
        n, _ = strconv.ParseInt(string(char), 10, 64)
        sum += n
       }
	fmt.Println(sum)
}
