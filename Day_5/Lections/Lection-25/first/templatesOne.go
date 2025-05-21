package main

import (
	"os"
	"log"
	"text/template"
)


type Part struct {
	Name string
	Count int
}

type Subscriber struct {
	Name	string
	Rate	float64
	Active	bool 
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}
func main(){
	_, err := os.Stdout.Write([]byte("hello\n"))
	check(err)

	executeTemplate("Dot is: {{.}}!\n", "ABC" )
	executeTemplate("Dot is: {{.}}!\n", 42.5 )
	executeTemplate("start {{if .}}Dot it true!{{end}} finish\n", true)
	executeTemplate("start {{if .}}Dot it true!{{end}} finish\n", false)
	
	tmpl := "Before: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After: {{.}}\n"
	executeTemplate(tmpl, []string{"do", "re", "mi"})
	
	tmpl = "Prices: \n{{range .}}${{.}}\n{{end}}"
	executeTemplate(tmpl, []float64{1.25, 0.99, 27})
	executeTemplate(tmpl, []float64{}) // Пустой вывод
	executeTemplate(tmpl, nil)		   // Пустой вывод

	tmpl = "Name: {{.Name}}\nCount: {{.Count}}\n"
	executeTemplate(tmpl, Part{Name: "Keys", Count: 5})
	executeTemplate(tmpl, Part{Name: "Cables", Count: 2})

	tmpl = "Name: {{.Name}}\n{{if .Active}}Rate: {{.Rate}}\n{{end}}"
	subscriber := Subscriber{Name: "Ivan Ivanov", Rate: 4.99, Active: true}
	executeTemplate(tmpl, subscriber)
	subscriber = Subscriber{Name: "Petr Petrov", Rate: 5.99, Active: false}
	executeTemplate(tmpl, subscriber)
}