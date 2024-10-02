package main

import (
	"fmt"
)

func main()  {
	
// Побитовые операции
	// & - bit and
	// | - bit or
	// ^ - bit xor
	// >> - сдвиг вправо
	// << - сдвиг влево

	a := 5  // bin: 0101
	b := 10 // bin: 1010
	c := a & b 
	d := a | b
	fmt.Println("c:", c) // 0
	fmt.Println("d:", d) // 15
	fmt.Println(" 5 ^ 3 =", 5 ^ 3) // 101 ^ 011: 6
	fmt.Println(" 4 ^ 7 =", 4 ^ 7) // 100 ^ 111: 3 ЖЛ

}