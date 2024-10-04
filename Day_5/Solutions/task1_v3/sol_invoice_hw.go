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
	"errors"
	"fmt"
	"regexp"
)

// Address представляет адрес доставки
type Address struct {
	PostalCode string
	City       string
	Street     string
	House      string
	Apartment  string
}

// Invoice представляет накладную
type Invoice struct {
	ProductName  string
	Quantity     int
	CustomerName string
	PhoneNumber  string
	DeliveryAddress Address
}

// NewInvoice создает новый объект Invoice с проверкой полей
func NewInvoice(productName string, quantity int, customerName, phoneNumber string, address Address) (*Invoice, error) {
	if err := validateFields(productName, quantity, customerName, phoneNumber, address); err != nil {
		return nil, err
	}
	return &Invoice{
		ProductName:  productName,
		Quantity:     quantity,
		CustomerName: customerName,
		PhoneNumber:  phoneNumber,
		DeliveryAddress: address,
	}, nil
}

// validateFields проверяет правильность заполнения полей
func validateFields(productName string, quantity int, customerName, phoneNumber string, address Address) error {
	if len(productName) == 0 || len(productName) > 100 {
		return errors.New("наименование товара должно быть от 1 до 100 символов")
	}
	if quantity <= 0 {
		return errors.New("количество должно быть положительным числом")
	}
	if !isValidName(customerName) {
		return errors.New("ФИО покупателя должно содержать только буквы")
	}
	if !isValidPhoneNumber(phoneNumber) {
		return errors.New("контактный телефон должен содержать ровно 10 цифр")
	}
	if !isValidAddress(address) {
		return errors.New("адрес должен содержать корректные данные")
	}
	return nil
}

// isValidName проверяет, что имя состоит только из букв
func isValidName(name string) bool {
	re := regexp.MustCompile(`^[a-zA-Zа-яА-ЯёЁ\s]+$`)
	return re.MatchString(name)
}

// isValidPhoneNumber проверяет номер телефона на наличие ровно 10 цифр
func isValidPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\d{10}$`)
	return re.MatchString(phone)
}

// isValidAddress проверяет корректность адреса
func isValidAddress(address Address) bool {
	re := regexp.MustCompile(`^\d{6}$`)
	return re.MatchString(address.PostalCode) && len(address.City) > 0 && len(address.Street) > 0 && len(address.House) > 0
}

// Print выводит информацию о накладной
func (i *Invoice) Print() {
	fmt.Printf("Накладная:\n")
	fmt.Printf("Наименование товара: %s\n", i.ProductName)
	fmt.Printf("Количество: %d\n", i.Quantity)
	fmt.Printf("ФИО покупателя: %s\n", i.CustomerName)
	fmt.Printf("Контактный телефон: %s\n", i.PhoneNumber)
	fmt.Printf("Адрес доставки: %s, %s, ул. %s, дом %s, квартира %s\n",
		i.DeliveryAddress.PostalCode,
		i.DeliveryAddress.City,
		i.DeliveryAddress.Street,
		i.DeliveryAddress.House,
		i.DeliveryAddress.Apartment)
}

func main() {
	address := Address{
		PostalCode: "123456",
		City:       "Москва",
		Street:     "Тверская",
		House:      "1",
		Apartment:  "101",
	}

	invoice, err := NewInvoice("Товар A", 2, "Иван Иванов", "1234567890", address)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	invoice.Print()
}