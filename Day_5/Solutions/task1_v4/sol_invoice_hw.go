/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100)
2) Количество (только числа)
3) ФИО покупателя (только буквы)
4) Контактный телефон (10 цифр)
5) Адрес: индекс(ровно 6 цифр), город, улица, дом, квартира

Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.

Реализовать конструктор и несколько методов у типа "Накладная"

Пример:
invoice = NewInvoice()

или

order = NewOrder()
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Address struct {
	postalCode string
	city       string
	street     string
	building   string
	flat       string
}

type Customer struct {
	fullName string
	phone    string
	address  *Address
}

type Invoice struct {
	itemName string
	quantity string
	customer *Customer
}

func newAddress(
	postalCode string,
	city string,
	street string,
	building string,
	flat string) *Address {
	return &Address{postalCode, city, street, building, flat}
}

func newCustomer(
	fullName string,
	phone string,
	address *Address) *Customer {
	return &Customer{fullName, phone, address}
}

func newInvoice(
	itemName string,
	quantity string,
	customer *Customer) *Invoice {
	return &Invoice{itemName, quantity, customer}
}

const exit string = "выход"

func main() {
	printDelimiter()
	fmt.Println("Для оформления заказа необходимо зарегистрироваться.")
	fmt.Println("Для прекращения процедуры оформления введите:", exit)
	printDelimiter()

	var success bool = true

	customer, success := readCustomer(success)
	printDelimiter()

	if success {
		fmt.Println("Вы успешно зарегистрированы. Теперь вы можете оформить заказ.")
		fmt.Println("Для прекращения процедуры оформления введите:", exit)
		printDelimiter()

		invoice, sucess := readInvoice(success, customer)
		printDelimiter()

		if sucess {
			fmt.Println("Товар успешно оформлен.")
			fmt.Println("Сформирована накладная")
			printDelimiter()
			invoice.PrintInvoice()
			printDelimiter()
		}
	}
}

func (invoice *Invoice) PrintInvoice() {
	fmt.Println("ФИО покупателя:", invoice.customer.fullName)
	fmt.Println("Телефон покупателя:", invoice.customer.phone)
	fmt.Printf("Адрес покупателя: %s, %s, %s, %s, %s\n",
		invoice.customer.address.postalCode,
		invoice.customer.address.city,
		invoice.customer.address.street,
		invoice.customer.address.building,
		invoice.customer.address.flat)
	fmt.Println("Наименование товара:", invoice.itemName)
	fmt.Println("Количество товара:", invoice.quantity)
}

func readCustomer(readFlag bool) (*Customer, bool) {

	fullname, readFlag := readParameter(
		readFlag,
		"Введите ФИО",
		"ФИО должно содержать не менее двух слов и состоять только из букв",
		isCustomerNameValid)

	phone, readFlag := readParameter(
		readFlag,
		"Введите номер телефона в формате XXXXXXXXXX (10 цифр)",
		"Номер телефона должен состоять только из цифр",
		isCustomerPhoneValid)

	address, readFlag := readAddress(readFlag)

	if readFlag {
		return newCustomer(fullname, phone, address), true
	}

	return nil, false
}

func readAddress(readFlag bool) (*Address, bool) {
	postalCode, readFlag := readParameter(
		readFlag,
		"Введите почтовый индекс XXXXXX (6 цифр)",
		"Почтовый индекс должен состоять только из цифр",
		isPostalCodeValid)
	city, readFlag := readParameter(
		readFlag,
		"Введите город (только буквы)",
		"Наименование города должно состоять только из букв",
		isOnlyLetters)
	street, readFlag := readParameter(
		readFlag,
		"Введите улицу (только буквы)",
		"Наименование улицы должно состоять только из букв",
		isOnlyLetters)
	building, readFlag := readParameter(
		readFlag,
		"Введите номер дома (только цифры)",
		"Номер дома должен состоять только из цифр",
		isOnlyDigits)
	flat, readFlag := readParameter(
		readFlag,
		"Введите номер квартиры (только цифры)",
		"Номер квартры должен состоять только из цифр",
		isOnlyDigits)

	if readFlag {
		return newAddress(postalCode, city, street, building, flat), true
	}

	return nil, false
}

func readInvoice(readFlag bool, customer *Customer) (*Invoice, bool) {
	itemName, readFlag := readParameter(
		readFlag,
		"Введите наименование товара (от 1 до 100 символов)",
		"Наименование товара должно состоять только из букв.",
		isItemNameValid)

	quantity, readFlag := readParameter(
		readFlag,
		"Введите количество (только цифры)",
		"Количество должно состоять только из цифр",
		isOnlyDigits)

	if readFlag {
		return newInvoice(itemName, quantity, customer), true
	}

	return nil, false
}

func readParameter(
	readFlag bool,
	question string,
	error string,
	check func(string) bool) (string, bool) {

	if !readFlag {
		return "", false
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s: ", question)

		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\n")

		if input == exit {
			break
		}

		if check(input) {
			return input, true
		}

		fmt.Printf("Ошибка: %s.\n", error)
	}

	return "", false
}

func isItemNameValid(fullName string) bool {
	r, _ := regexp.Compile(`^[А-Яа-я]{1,100}$`)

	return r.MatchString(fullName)
}

func isCustomerNameValid(fullName string) bool {
	r, _ := regexp.Compile(`^[А-Яа-я]+\s[А-Яа-я]+\s*[А-Яа-я]*$`)

	return r.MatchString(fullName)
}

func isCustomerPhoneValid(number string) bool {
	r, _ := regexp.Compile(`^\d{10}$`)

	return r.MatchString(number)
}

func isPostalCodeValid(number string) bool {
	r, _ := regexp.Compile(`^\d{6}$`)

	return r.MatchString(number)
}

func isOnlyLetters(input string) bool {
	r, _ := regexp.Compile(`^[А-Яа-я]+$`)

	return r.MatchString(input)
}

func isOnlyDigits(input string) bool {
	r, _ := regexp.Compile(`^\d+$`)

	return r.MatchString(input)
}

func printDelimiter() {
	fmt.Println(strings.Repeat("=", 55))
}
