package app

import (
	"context"

	"github.com/imirjar/rb-michman/config"
	"github.com/imirjar/rb-michman/internal/gateway/http"
	"github.com/imirjar/rb-michman/internal/service/grazer"
	"github.com/imirjar/rb-michman/internal/service/reporter"
	"github.com/imirjar/rb-michman/internal/storage/collector"
	"github.com/imirjar/rb-michman/internal/storage/diver"
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
	memStore := collector.New()

	grazer := grazer.New()
	reporter := reporter.New()

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
		srv.Start(ctx, cfg.GetMichmanAddr())
		done <- true
	}()

	<-done // Ожидание завершения первой горутины
	return nil
}