/* 
    Задача №1
    Написать функцию, которая расшифрует строку. 
    code = "220411112603141304"
    Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
    Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
    Вход: строка из цифр. Выход: Текст.
    Проверка работы функции выполняется через другую строку.

    Задача №1.1
    Реализовать и функцию зашифровки

    codeToString(code) -> "???????' 

    stringToCode("hello") -> "??????"
*/

package main

import (
    
    "fmt"
    "strconv"
    "errors"
)

func main() {
    input := "220411112603141304"
    fmt.Println("input:", input)
    decoded, _ := decode(input)
    fmt.Println("decoded:", decoded)
    encoded, _ := encode(decoded)
    fmt.Println("encoded:", encoded)
    fmt.Println("Are equal:", input == encoded)

	encoded_go, _ := encode("hello golang")
    fmt.Println("encoded_go:", encoded_go)
}

func decode(input string) (string, error) {
    if len(input) < 2 || len(input) % 2 != 0 {
        return "", errors.New("Wrong input")
    }
    
    arr := make([]byte, len(input) / 2)
    
    for i := 0; i < len(input); i +=2 {
        var stringCode = input[i:i+2]
        var code, _ = strconv.Atoi(stringCode)
        var convertedCode, convertError = convert(byte(code))
        
        if convertError != nil {
            return "", errors.New("Wrong input")
        }
        
        arr[ i / 2 ] = convertedCode
    }
    
    return string(arr), nil
}

func encode(input string) (string, error) {
    var result string
    
    for i := 0; i < len(input); i++ {
        var symbol = input[i]
        var code = byte(symbol)
        var encodedCode, convertError = convertBack(code)
        
        if convertError != nil {
            return "", errors.New("Wrong input")
        }
        
        result += fmt.Sprintf("%02d", encodedCode)
    }
    
    return result, nil
}

func convert(input byte) (byte, error) {
    if input < 0 || input > 26 {
        return 0, errors.New("Wrong input")
    }
    
    if input != 26 {
       return input + 97, nil
    } else {
        return 32 , nil
    }
}

func convertBack(input byte) (byte, error) {
    if !(input == 32 || (input >= 97 && input <= 122)) {
        return 0, errors.New("Wrong input")
    }
    
    if input != 32 {
       return input - 97, nil
    } else {
        return 26 , nil
    }
}