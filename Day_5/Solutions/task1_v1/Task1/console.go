package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const defaultTriesCount = 10

func GetCustomer() (*Person, error) {
	fmt.Println("Введите ФИО покупателя в формате Фамилия Имя Отчество(если есть)")
	triesCount := 0
	for triesCount < defaultTriesCount {
		customerInfo := ""
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			customerInfo = scanner.Text()
		}

		customerFIO := strings.Split(customerInfo, " ")
		if len(customerFIO) < 3 {
			triesCount++
			fmt.Println("Введите корректные ФИО покупателя.")
			continue
		}

		customer := NewPerson(customerFIO[0], customerFIO[1], customerFIO[2])
		err := customer.Validate()
		if err == nil {
			return customer, nil
		}

		triesCount++
		fmt.Println("Введите корректные ФИО покупателя:", err)
	}
	fmt.Printf("Количество попыток ввода превысило %d.\n", defaultTriesCount)
	return nil, errors.New("превышено количество попыток создания пользователя")
}

func GetAddress() (*Address, error) {
	address := NewEmptyAddress()
	errIndex := SetIndex(address)
	if errIndex != nil {
		return nil, errIndex
	}

	errCity := SetCity(address)
	if errCity != nil {
		return nil, errCity
	}

	// TODO:дальнейшую валидацию нужно сделать по аналогии как выше
	fmt.Println("Введите улицу, дом и квартиру")
	addressInfo := ""
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		addressInfo = scanner.Text()
	}
	addressInfoSplit := strings.Split(addressInfo, " ")
	address.Street = addressInfoSplit[0]
	address.House = addressInfoSplit[1]
	address.Apartment = addressInfoSplit[2]

	return address, nil
}

func SetIndex(address *Address) error {
	fmt.Println("Введите индекс (6 цифр)")
	triesCount := 0
	for triesCount < defaultTriesCount {
		indexStr := ""
		fmt.Scanln(&indexStr)

		err := address.SetIndex(indexStr)
		if err == nil {
			return nil
		}

		triesCount++
		fmt.Println(err)
		fmt.Println("Введите корректный индекс из 6 цифр")
	}

	fmt.Printf("Количество попыток ввода превысило %d.\n", defaultTriesCount)
	return errors.New("превышено количество попыток установки индекса")
}

func SetCity(address *Address) error {
	fmt.Println("Введите город")
	triesCount := 0
	for triesCount < defaultTriesCount {
		cityStr := ""
		fmt.Scanln(&cityStr)

		err := address.SetCity(cityStr)
		if err == nil {
			return nil
		}
		triesCount++
		fmt.Println("Введите корректное название города:", err)
	}

	fmt.Printf("Количество попыток ввода превысило %d.\n", defaultTriesCount)
	return errors.New("превышено количество попыток установки индекса")
}
