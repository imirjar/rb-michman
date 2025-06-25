package app

import (
	"context"
	"log"

	"github.com/imirjar/rb-michman/config"
	"github.com/imirjar/rb-michman/internal/gateway/http"
	"github.com/imirjar/rb-michman/internal/service"
	"github.com/imirjar/rb-michman/internal/storage/queries"
	"github.com/imirjar/rb-michman/internal/storage/reports"
)

func Run(ctx context.Context) error {
	// Подгружаем файл конфигураций /config
	cfg := config.New()

	// Приложение состоит из 3 основных слоев
	// 1)srv - Шлюз через который приложений обрабатывает входящие запросы
	// 2)service - Сервис для логики обработки запросов
	// 3)storage - Хранилище значений

	// Слои представляют собой цепочку последовательностей
	// Соединяем srv->service->storage
	// Так данные и будут двигаться

	// Report storage for michman and diver
	reportStore := reports.New()
	reportStore.Connect(context.Background(), cfg.Mongo)
	if err := reportStore.Connect(context.Background(), cfg.Mongo); err != nil {
		log.Print(err)
	}
	defer reportStore.Disconnect()

	// Quries from michman to diver
	queryStore := queries.New()
	if err := queryStore.Connect(context.Background(), cfg.Rabbit); err != nil {
		log.Print(err)
	}
	defer queryStore.Disconnect()

	// Main business logic
	service := service.New()
	service.MQ = queryStore
	service.RS = reportStore

	srv := http.New()
	srv.Service = service

	done := make(chan bool)

	//Запускаем сервер
	go func() {
		if err := srv.Start(ctx, cfg.Port); err != nil {
			log.Print(err)
			done <- true
		}
	}()

	<-done // Ожидание завершения первой горутины
	log.Print("shutdown")
	return nil
}
