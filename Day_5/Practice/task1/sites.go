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


// func responseSize(...) { ... }
	

// func main() {...}
