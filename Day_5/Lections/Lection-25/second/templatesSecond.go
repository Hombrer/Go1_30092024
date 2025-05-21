package main

import (
	"log"
	"html/template"
	"os"
)

type Invoice struct {
	Name	string
	Paid	bool
	Charges	[]float64
	Total	float64
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main(){
	html, err := template.ParseFiles("bill.html")
	check(err)
	bill := Invoice{
		Name:	 "Ivan Petrov",
		Paid:	 true,
		Charges: []float64{23.19, 1.13, 42.79},
		Total:	 67.11,
	}
	err = html.Execute(os.Stdout, bill)
	check(err)
}