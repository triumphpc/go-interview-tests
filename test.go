package main

import (
	"fmt"
	"log"
	"time"
)

type Order struct {
	ID   int
	Data string
}

var Results map[Order]string // Не инициировано

func processOrders() {
	orders := make(chan Order)   // канал для заказов
	results := make(chan string) // канал для результатов

	go func() {
		for order := range orders {
			result := fmt.Sprintf("Заказ %d обработан", order.ID)
			results <- result

			time.Sleep(1 * time.Second)

		}
		close(results)

	}()

	go func() {
		for order := range results {
			log.Println(order)
		}
	}()

	// Основной цикл
	for i := 1; i <= 3; i++ {
		orders <- Order{ID: i, Data: "данные"}
	}

	log.Println("HERE 1")

	close(orders)

	log.Println("HERE 2")

}

func main() {
	processOrders()
}
