package main

import (
	"fmt"
	"regexp"
)

const (
	MinCityLength      = 2
	MaxCityLength      = 100
	MinStreetLength    = 4
	MaxStreetLength    = 100
	MinHouseLength     = 1
	MaxHouseLength     = 10
	MaxApartmentLength = 5
)

type Address struct {
	Index     string
	City      string
	Street    string
	House     string
	Apartment string
}

func NewEmptyAddress() *Address {
	return &Address{}
}

func NewAddress(index string, city string, street string, house string, apartment string) *Address {
	return &Address{
		Index:     index,
		City:      city,
		Street:    street,
		House:     house,
		Apartment: apartment,
	}
}

func (a *Address) SetIndex(index string) error {
	re := regexp.MustCompile(`^\d{6}$`)
	if len(index) != 6 && !re.MatchString(index) {
		return fmt.Errorf("индекс должен состоять из 6 цифр")
	}

	a.Index = index
	return nil
}

func (a *Address) SetCity(city string) error {
	if len(city) <= MinCityLength || len(city) >= MaxCityLength {
		return fmt.Errorf("название города должно быть длиной от %d до %d символов",
			MinCityLength, MaxCityLength)
	}

	re := regexp.MustCompile(`^[a-zA-Zа-яА-Я0-9\s\-]+$`)
	if !re.MatchString(city) {
		return fmt.Errorf("в названии города могут быть только буквы, цифры, тире и пробелы")
	}

	a.City = city
	return nil
}

func (a *Address) SetStreet(street string) error {
	if len(street) <= MinStreetLength || len(street) >= MaxStreetLength {
		return fmt.Errorf("название улицы должно быть длиной от %d до %d",
			MinStreetLength, MaxStreetLength)
	}

	re := regexp.MustCompile(`^[a-zA-Zа-яА-Я0-9\s\.\-]+$`)
	if !re.MatchString(street) {
		return fmt.Errorf("в названии улицы могут быть только буквы, цифры, пробелы, тире и точки")
	}

	a.Street = street
	return nil
}

func (a *Address) SetHouse(house string) error {
	if len(house) <= MinHouseLength || len(house) >= MaxHouseLength {
		return fmt.Errorf("название дома должно быть длиной от %d до %d",
			MinHouseLength, MaxHouseLength)
	}

	a.House = house
	return nil
}

func (a *Address) SetApartment(apartment string) error {
	if len(apartment) >= MaxApartmentLength {
		return fmt.Errorf("номер квартиры должно быть длиной до %d", MaxApartmentLength)
	}

	a.Apartment = apartment
	return nil
}
