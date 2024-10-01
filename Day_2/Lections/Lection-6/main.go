package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	// Numerics. Integers
	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64
	var a int = 32
	b := 92
	fmt.Println("Value of a:", a, "Value of b:", b, "Sum of a+b:", a + b)

	// Тип данных можно получить через %T форматирование
	fmt.Printf("type of a: %T and type of b: %T\n", a, b)

	// Узнаем, сколько байт в памити занимает тип int
	// При использовании короткого объявления int -> платформо-зависимый
	fmt.Printf("Type %T sizes %d bytes in memory.\n", a, unsafe.Sizeof(a))
	fmt.Printf("Type %T sizes %d bytes in memory.\n", b, unsafe.Sizeof(b))

	// Если выполняются арифметические операции над int и intX,
	// то обязательно нужно использовать механизм приведения типов, 
	// т.к. int != int64
	var c64 int64 = 16_213_897
	var usualInt int = 123_567
	fmt.Println(c64 + int64(usualInt))

	// Аналогичным образом устроены uint8, uint16, uint32, uint64

	// Numerics. Float
	// float32, float64
	floatFirst, floatSecond := 5.67, 12.54
	fmt.Printf("Type %T sizes %d bytes in memory.\n", floatFirst, unsafe.Sizeof(floatFirst))
	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond
	fmt.Printf("Sum: %.3f, Sub: %.3f\n", sum, sub)
	fmt.Println(0.1 * 3 == 0.3)
	fmt.Println(4.2 - 3 == 1.2)
	fmt.Println(math.Pow(floatFirst, floatSecond))

	// Numerics. Complex
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

	// Strings. Строка - это набор байт
	name := "Федя"
	lastname := "Pupkin"
	fi := name + " " + lastname
	fmt.Println("Full name:", fi)
	// Функция len() - кол-во байт в строке
	fmt.Println("Length of string", len(name))
	fmt.Println("Length of string", len(lastname))

	// Rune - руна. Это один utf-ный символ
	fmt.Println("Кол-во символов в строке:", utf8.RuneCountInString(name))

	// Поиск подстроки в строке
	totalString, subString := "ABCDEFG", "CDE"
	fmt.Println(strings.Contains(totalString, subString))

	// rune -> это alias int32
	var sampleRune rune
	var anotherRune = 'Q' // Для инициализации руны используйте ''
	var thirdRune = 234
	fmt.Println(sampleRune)
	fmt.Printf("Rune as char: %c\n", sampleRune)
	fmt.Println(anotherRune)
	fmt.Printf("Rune as char: %c\n", anotherRune)
	fmt.Println(thirdRune)
	fmt.Printf("Rune as char: %c\n", thirdRune)

	// -1 if first < second, 0 if first == second, 1 if first > second
	fmt.Println(strings.Compare("abcd", string('a')))

	var aByte byte // alias uint8
	fmt.Println("Byte", aByte)
}