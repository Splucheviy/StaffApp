package app

import (
	"context"
	"log"
)

// App представляет основную структуру приложения
type App struct {
	logger *log.Logger
}

// New создает новый экземпляр приложения
func New(logger *log.Logger) *App {
	return &App{
		logger: logger,
	}
}

// Run запускает приложение
func (a *App) Run(ctx context.Context) error {
	a.logger.Println("Starting application...")

	// TODO: Добавить инициализацию компонентов
	// - Инициализация базы данных
	// - Инициализация репозиториев
	// - Инициализация use cases
	// - Инициализация HTTP сервера

	return nil
}
