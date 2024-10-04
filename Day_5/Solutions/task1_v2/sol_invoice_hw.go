/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100) ++
2) Количество (только числа) ++
3) ФИО покупателя (только буквы) ++
4) Контактный телефон (10 цифр) ++
5) Адрес: индекс(ровно 6 цифр), город, улица, дом, квартира ++

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
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"regexp"
)

type Order struct {
	product    Product
	buyer 	   Buyer
}

type Product struct {
	nameAndCount map[string]int64
}

type Buyer struct {
	fullName   Fullname
	phone 	   int64
	address    Address
}

type Fullname struct {
	firstName  string
	lastName   string
	patronymic string
}

type Address struct {
	index 	   int64
	city 	   string
	street	   string
	house	   string
	flat  	   string
}

func main() {
	ord := NewOrder()
	fmt.Println("Номер телефона из заказа:",ord.GetPhone())
	fmt.Println("Город из заказа:",ord.GetCity())
	fmt.Printf("Type: %T and value %v\n", ord, ord)
}

func NewOrder() *Order {
	return &Order{
		product: *NewProduct(),
		buyer: *NewBuyer(), 	   
    }
}

func NewProduct() *Product {
	start: 
	count := inputInt("Введите количество товаров планируемых к покупке, но не больше 100")
	nameAndCount := make(map[string]int64)
	if count > 0 && count <= 100 {
		for i := 0; i < int(count); i++ {
			nameAndCount[inputStr("Введите название товара")] = inputInt("Введите количество товара")
		}
	} else {
		fmt.Println("Ошибка при вводе количества товара, введите количество товара заного")
		goto start
	}
	return &Product{
		nameAndCount: nameAndCount,
	}
}

func NewBuyer() *Buyer {
	start:
	phone := inputInt("Введите телефон")
    re := regexp.MustCompile(`^((\+7|7|8)+([0-9]){10})$`)
	strphone := strconv.Itoa(int(phone))
    match := re.MatchString(strphone)
    if match {
		return &Buyer{
			fullName: *NewFullname(),
			phone: phone,
			address: *NewAddress(),
		}
    } else {
        fmt.Println("Телефонный номер введен не корректно введите его заного")
		goto start
	}
}

func NewFullname() *Fullname {
	start:
	firstName := inputStr("Введите Имя")
	lastName := inputStr("Введите Фамилию")
	patronymic := inputStr("Введите Отчество")
    re := regexp.MustCompile(`^[а-яА-ЯёЁa-zA-Z]+$`)
    match := re.MatchString(firstName+lastName+patronymic)
    if match && firstName != "" && lastName != "" && patronymic != "" {
		return &Fullname{
			firstName: firstName,
			lastName: lastName,
			patronymic: patronymic,
		}
    } else {
        fmt.Println("ФИО содержит не только буквы или пусты, введите ФИО заного")
		goto start
	}
}

func NewAddress() *Address{
	start:
	index := inputInt("Введите Индекс")
	indexStr := strconv.Itoa(int(index))
	if len(indexStr) == 6 {
		return &Address{
			index:  index,
			city:   inputStr("Введите Город"),
			street: inputStr("Введите Улицу"),
			house:	inputStr("Введите дом"),
			flat:  	inputStr("Введите квартиру"),
		}
    } else {
        fmt.Println("Индекс введен не верпно должно быть 6 знаков, введите индекс заного")
		goto start
	}
}

func (o Order) GetPhone() int64 {
	return o.buyer.phone
}

func (o Order) GetCity() string {
	return o.buyer.address.city
}

func inputStr(text string) string {
	fmt.Println(text)
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	return str
}

func inputInt(text string) int64 {
	start:
	fmt.Println(text)
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	count, err := strconv.ParseInt(strings.TrimSuffix(str, "\n"), 10, 64)
	if err != nil {
		fmt.Println("Можно вводить только цифры") 
		goto start
	}
	return count
}