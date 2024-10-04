package main

import "fmt"

type Invoice struct {
	Goods    []string
	Count    int
	Customer Person
	Phone    string
	Address  Address
}

func NewInvoice(Customer Person, Phone string, Address Address) *Invoice {
	return &Invoice{
		Goods:    []string{},
		Customer: Customer,
		Phone:    Phone,
		Address:  Address,
	}
}

func (i *Invoice) AddGoods(goods ...string) {
	i.Goods = append(i.Goods, goods...)
	i.Count += len(goods)
}

func (i *Invoice) Print() {
	fmt.Println("Товары:", i.Goods)
	fmt.Println("Кол-во:", i.Count)
	fmt.Println("Покупатель:", i.Customer)
	fmt.Println("Телефон:", i.Phone)
	fmt.Println("Адрес:", i.Address)
}
