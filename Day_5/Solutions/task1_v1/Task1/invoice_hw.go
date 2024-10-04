package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Заполните накладную:")

	customer, err := GetCustomer()
	if err != nil {
		return
	}

	address, err := GetAddress()
	if err != nil {
		return
	}

	phone := GetPhone()
	invoice := NewInvoice(*customer, phone, *address)

	//TODO:вынести в отдельный метод, добавить валидацию товаров
	fmt.Println("Напишите список товаров через пробел.")
	var goodsStr string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		goodsStr = scanner.Text()
	}
	goods := strings.Split(goodsStr, " ")
	invoice.AddGoods(goods...)
	fmt.Println("Ваша накладная:")
	//TODO:оформить вывод покупателя, адреса и товаров более user-friendly
	invoice.Print()
}

// TODO:не доделано, нужна возможность повторного ввода
func GetPhone() string {
	var phone string
	fmt.Println("Введите номер телефона:")
	fmt.Scan(&phone)
	errPhone := ValidatePhoneNumber(phone)
	if errPhone != nil {
		fmt.Println("Ошибка:", errPhone)
		return ""
	}
	return phone
}

func ValidatePhoneNumber(phone string) error {
	phoneRegex := regexp.MustCompile(`^[0-9]+$`)

	if !phoneRegex.MatchString(phone) {
		return errors.New("в телефоне должны быть только цифры")
	}
	return nil
}
