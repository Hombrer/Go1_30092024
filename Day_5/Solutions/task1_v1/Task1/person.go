package main

import (
	"fmt"
	"regexp"
)

type Person struct {
	FirstName string
	LastName  string
	Surname   string
}

func NewPerson(firstName string, lastName string, surName string) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Surname:   surName,
	}
}

func (p *Person) Validate() error {
	if len(p.FirstName) == 0 {
		return fmt.Errorf("имя не должно быть пустым")
	}

	if len(p.LastName) == 0 {
		return fmt.Errorf("фамилия не должна быть пустой")
	}

	re := regexp.MustCompile(`^[a-zA-Zа-яА-Я\-]+$`)
	if !re.MatchString(p.FirstName) {
		return fmt.Errorf("имя должно состоять из букв")
	}

	if !re.MatchString(p.LastName) {
		return fmt.Errorf("фамилия должна состоять из букв")
	}

	return nil
}
