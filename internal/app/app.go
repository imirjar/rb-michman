package app

import (
	"context"
	"log"

	"github.com/imirjar/rb-michman/config"
	"github.com/imirjar/rb-michman/internal/gateway/http"
	diver1 "github.com/imirjar/rb-michman/internal/service/diver"
	"github.com/imirjar/rb-michman/internal/service/grazer"
	"github.com/imirjar/rb-michman/internal/storage/diver"
	grazer1 "github.com/imirjar/rb-michman/internal/storage/grazer"
)

func Run(ctx context.Context) error {
	// Подгружаем файл конфигураций /config
	cfg := config.New()

	// Приложение состоит из 3 основных слоев
	// 1)rpc - Шлюз через который приложений обрабатывает входящие запросы
	// 2)service - Сервис для логики обработки запросов
	// 3)storage - Хранилище значений
	// Слои представляют собой цепочку последовательностей
	apiStore := diver.New()
	memStore := grazer1.New()

	// Collect information about current divers
	grazer := grazer.New()

	// Current Divers API which are available
	reporter := diver1.New()

	srv := http.New()

	// Соединяем srv->service->storage
	// Так данные и будут двигаться
	grazer.Collector = memStore
	grazer.Divers = apiStore

	reporter.Divers = apiStore
	reporter.Collector = memStore

	srv.GrazerService = grazer
	srv.ReportService = reporter

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
