package main

// 3.2 Доработать предыдущую задачу так, чтобы выводились только числа,
//     делящиеся на 5 без остатка.

import (
    "fmt"
    "math/big"
)

func main() {
    var (
        a float64
        b float64
    )
    println("Введите два числа a и b (b больше a)")
    fmt.Scan(&a, &b)
    for i := int32(a+1); i < int32(b); i++ {
        if i % 5 == 0 {
            fmt.Printf("%v ", i)
        }
    }
    if big.NewFloat(b).Cmp(big.NewFloat(float64(int32(b)))) == 1 && int32(b) % 5 == 0 {
        println(int32(b))
    }
}