package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Инициализация логгера
	logger := log.New(os.Stdout, "[STAFF-APP] ", log.LstdFlags|log.Lshortfile)

	// Загрузка конфигурации
	// TODO: Добавить загрузку конфигурации

	// Инициализация приложения
	// TODO: Добавить инициализацию приложения

	// Обработка сигналов завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	logger.Println("Shutting down server...")
}
