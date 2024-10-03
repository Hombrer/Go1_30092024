package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const maxTries = 5
	fmt.Println(`Hello!
Password length must be between 8 and 15 characters
and contain at least one digit, one lowercase letter,
one uppercase letter and one special character: _!@#$%^&`)

	for i := 0; i < maxTries; i++ {
		fmt.Println("Password attempt number: " + strconv.Itoa(i+1))
		fmt.Println("Enter password:")
		var password string
		fmt.Scan(&password)

		isSuccess, errors := checkPassword(password)
		if isSuccess {
			fmt.Println("Password is valid")
			return
		}
		fmt.Println(errors)
	}

	fmt.Println("No attempts left.")
}

const (
	Digits = iota
	Lowercase
	Uppercase
	Specials
)

func checkPassword(password string) (isSuccess bool, errors string) {
	const minLength = 8
	const maxLength = 15
	if len(password) < minLength || len(password) > maxLength {
		message := fmt.Sprintf("Password length must be between %d and %d characters", minLength, maxLength)
		return false, message
	}

	requiredCharacters := map[int]string{
		Digits:    "digits",
		Lowercase: "lowercase English letters",
		Uppercase: "UPPERCASE English letters",
		Specials:  "special characters: _!@#$%^&",
	}

	for _, character := range password {
		switch {
		case isSpecialCharacter(character):
			delete(requiredCharacters, Specials)
		case isDigit(character):
			delete(requiredCharacters, Digits)
		case isLowercase(character):
			delete(requiredCharacters, Lowercase)
		case isUppercase(character):
			delete(requiredCharacters, Uppercase)
		}

		if len(requiredCharacters) == 0 {
			return true, ""
		}
	}

	errors = "Password must contain: "
	for _, errorMessage := range requiredCharacters {
		errors += errorMessage + ", "
	}
	errors = strings.TrimSuffix(errors, ", ") + "."
	return false, errors
}

func isSpecialCharacter(character rune) bool {
	return strings.ContainsRune("_!@#$%^&", character)
}

func isDigit(character rune) bool {
	return character >= '0' && character <= '9'
}

func isLowercase(character rune) bool {
	return character >= 'a' && character <= 'z'
}

func isUppercase(character rune) bool {
	return character >= 'A' && character <= 'Z'
}
