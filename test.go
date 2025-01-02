package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

// Постановка задачи:
// Уведомления поступают из различных источников и могут иметь разные форматы (например, CSV и JSON).
// Уведомления должны быть отправлены в различные каналы (например, email, Telegram).
// Система должна быть легко расширяемой для добавления новых форматов источников и новых каналов отправки.
// Обработка уведомлений должна быть выполнена параллельно для повышения производительности.

// NotificationSource определяет интерфейс для источников уведомлений
type NotificationSource interface {
	Parse(data string) Notification
}

// NotificationChannel определяет интерфейс для каналов отправки уведомлений
type NotificationChannel interface {
	Send(notification Notification)
}

// Notification представляет структуру уведомления
type Notification struct {
	Name    string
	Type    string
	ID      string
	Content string
}

// CSVSource реализует интерфейс NotificationSource для CSV данных
type CSVSource struct{}

func (c CSVSource) Parse(data string) Notification {
	reader := csv.NewReader(strings.NewReader(data))
	record, _ := reader.Read()
	return Notification{
		Name:    record[0],
		Type:    record[1],
		ID:      record[2],
		Content: record[3],
	}
}

// JSONSource реализует интерфейс NotificationSource для JSON данных
type JSONSource struct{}

func (j JSONSource) Parse(data string) Notification {
	var notification Notification
	json.Unmarshal([]byte(data), &notification)
	return notification
}

// EmailChannel реализует интерфейс NotificationChannel для отправки email
type EmailChannel struct{}

func (e EmailChannel) Send(notification Notification) {
	fmt.Printf("Sending email to %s with content: %s\n", notification.ID, notification.Content)
}

// TelegramChannel реализует интерфейс NotificationChannel для отправки в Telegram
type TelegramChannel struct{}

func (t TelegramChannel) Send(notification Notification) {
	fmt.Printf("Sending telegram to %s with content: %s\n", notification.ID, notification.Content)
}

// Sender определяет канал отправки уведомлений на основе типа
type Sender struct {
	channels map[string]NotificationChannel
}

func (s *Sender) Send(notification Notification) {
	if channel, exists := s.channels[notification.Type]; exists {
		channel.Send(notification)
	} else {
		fmt.Printf("No channel found for type: %s\n", notification.Type)
	}
}

func main() {
	var wg sync.WaitGroup

	// Примеры данных
	csvData := "Jhon,email,name@example.com,some content"
	jsonData := `{"name":"Alice", "type": "telegram", "id":"@alice", "content":"some content"}`

	// Создаем каналы
	emailChannel := EmailChannel{}
	telegramChannel := TelegramChannel{}

	// Создаем Sender с мапой каналов
	sender := Sender{
		channels: map[string]NotificationChannel{
			"email":    emailChannel,
			"telegram": telegramChannel,
		},
	}

	// Данные для обработки с указанием типа источника
	dataSources := []struct {
		data       string
		sourceType string
	}{
		{csvData, "csv"},
		{jsonData, "json"},
	}

	// Обрабатываем данные из разных источников в параллельных горутинах
	for _, ds := range dataSources {
		wg.Add(1)
		go func(data string, sourceType string) {
			defer wg.Done()

			var source NotificationSource

			// Выбираем источник на основе типа
			switch sourceType {
			case "csv":
				source = CSVSource{}
			case "json":
				source = JSONSource{}
			default:
				fmt.Println("Unknown source type")
				return
			}

			notification := source.Parse(data)
			sender.Send(notification)
		}(ds.data, ds.sourceType)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
}
