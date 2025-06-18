package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Note struct {
	TimeOffset time.Duration
	Text       string
}

func main() {
	// Запрашиваем ФИО в начале
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите ФИО интервьюируемого: ")
	fio, _ := reader.ReadString('\n')
	fio = fio[:len(fio)-1] // Убираем символ новой строки

	var notes []Note
	startTime := time.Now()

	// Запускаем параллельный таймер для уведомлений каждые 5 минут
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				elapsed := time.Since(startTime)
				minutes := int(elapsed.Minutes())
				seconds := int(elapsed.Seconds()) % 60
				fmt.Printf("\n[Уведомление] Прошло времени: %02d:%02d\n> ", minutes, seconds)
			}
		}
	}()

	fmt.Println("\nСессия начата. Вводите заметки (для завершения введите 'exit'):")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		if text == "exit" {
			break
		}

		notes = append(notes, Note{
			TimeOffset: time.Since(startTime),
			Text:       text,
		})
	}

	// Выводим полный отчёт
	fmt.Printf("\n--- Отчёт по собеседованию ---\n")
	fmt.Printf("Кандидат: %s\n", fio)
	fmt.Printf("Начало сессии: %s\n", startTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("Длительность: %s\n\n", time.Since(startTime).Round(time.Second))

	fmt.Println("Хронология заметок:")
	for _, note := range notes {
		minutes := int(note.TimeOffset.Minutes())
		seconds := int(note.TimeOffset.Seconds()) % 60
		fmt.Printf("[%02d:%02d] %s\n", minutes, seconds, note.Text)
	}
}
