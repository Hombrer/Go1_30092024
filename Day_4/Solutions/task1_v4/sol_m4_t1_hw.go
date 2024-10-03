package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	alphabetEncryptMap := getAlphabetEncryptMap()
	alphabetDecryptMap := getAlphabetDecryptMap()

	fmt.Println("Enter the string to encrypt:")
	stringForEncrypt := ""
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		stringForEncrypt = scanner.Text()
	}

	encryptedString := stringToCode(stringForEncrypt, alphabetEncryptMap)
	fmt.Println("Encrypted string:", encryptedString)
	fmt.Println("Decrypted again string:", codeToString(encryptedString, 2, alphabetDecryptMap))

	fmt.Println("Enter the string to decrypt:")
	stringForDecrypt := ""
	if scanner.Scan() {
		stringForDecrypt = scanner.Text()
	}

	fmt.Println("Decrypted string:", codeToString(stringForDecrypt, 2, alphabetDecryptMap))
}

func getAlphabetEncryptMap() map[rune]string {
	const alphabetCount = 26
	alphabetMap := make(map[rune]string, alphabetCount+1)

	for i := 0; i < alphabetCount; i++ {
		letterCipher := fmt.Sprintf("%02d", i)
		alphabetMap[rune('a'+i)] = letterCipher
	}
	alphabetMap[' '] = strconv.Itoa(alphabetCount)

	return alphabetMap
}

func getAlphabetDecryptMap() map[string]rune {
	const alphabetCount = 26
	alphabetMap := make(map[string]rune, alphabetCount+1)

	for i := 0; i < alphabetCount; i++ {
		letterCipher := fmt.Sprintf("%02d", i)
		alphabetMap[letterCipher] = rune('a' + i)
	}
	alphabetMap[strconv.Itoa(alphabetCount)] = ' '

	return alphabetMap
}

func stringToCode(text string, ciphers map[rune]string) string {
	var codeBuilder strings.Builder
	for _, letter := range text {
		if value, ok := ciphers[letter]; ok {
			codeBuilder.WriteString(value)
		}
	}
	return codeBuilder.String()
}

func codeToString(code string, codeLength int, ciphers map[string]rune) string {
	decryptedLength := utf8.RuneCountInString(code) / codeLength
	decrypted := make([]rune, decryptedLength)
	for i := 0; i < utf8.RuneCountInString(code); i += codeLength {
		pair := code[i : i+codeLength]
		if value, ok := ciphers[pair]; ok {
			decrypted[i/codeLength] = value
		}
	}
	return string(decrypted)
}
