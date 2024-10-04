/*
Реализовать загрузку страниц с пяти разных сайтов с помощью горутин
Вывести в консоль URL сайта и размер первоначальной страницы в байтах

Подсказки, что нужно использовать:

import (

	"net/http"
	"io"

)

// Получить URL:
response, err := http.Get("www.example.com")

defer response.Body.Close()

// Получить тело ответа
body, err := io.ReadAll(response.Body)

// Рекомендую создать соответствующую структуру Page с полями URL и Size
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, pagesChan chan Page) {
	// Получить URL:
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	// Получить тело ответа
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	pagesChan <- Page{URL: url, Size: len(body)}
}

func main() {
	pagesChan := make(chan Page)

	pages := []string{
		"https://github.com",
		"https://yandex.ru",
		"https://google.com",
		"https://yahoo.com",
		"https://habr.com",
	}

	for _, page := range pages {
		go responseSize(page, pagesChan)
	}

	for i := 0; i < len(pages); i++ {
		page := <- pagesChan
		fmt.Println("URL:", page.URL, "Size:", page.Size)
	}
}
