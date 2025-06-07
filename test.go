package main

import "fmt"

type Order struct {
	ID   int
	Data string
}

var Results map[Order]string // Не инициировано

func processOrders() {
	orders := make(chan Order) // канал для заказов
	//var results chan string // 1. Не иницирована
	results := make(chan string) // канал для результатов

	// Горутина для обработки заказов
	go func() {
		for order := range orders {
			// Обработка заказа
			result := fmt.Sprintf("Заказ %d обработан", order.ID)
			results <- result

			Results[order] = order.Data // 4. Тут гонка данных
		}
	}()

	// Основной цикл
	for i := 1; i <= 3; i++ { // 2. нужно выносить во вторую горутину
		orders <- Order{ID: i, Data: "данные"}
		// fmt.Println(<-results) // 1. будет deadlock
	}

	close(orders)
	close(results) // 3. будет паника, если третья первая горутина не успеет все записать
}

func main() {
	processOrders()
}
