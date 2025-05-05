package main

import (
	"TestTask/handler"
	"TestTask/migration"
	"TestTask/repository"
	"TestTask/router"

	"log"
)

func main() {
	// 1. Инициализация базы данных и миграция
	migration.InitDB()

	// 2. Создание слоя репозитория (работа с БД)
	personRepo := repository.NewPersonRepository(migration.DB)

	// 3. Создание обработчиков (handlers)
	personHandler := handler.NewPersonHandler(personRepo)

	// 4. Настройка роутера
	r := router.SetupRouter(personHandler)

	// 5. Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
