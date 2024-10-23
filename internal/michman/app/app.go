package app

import (
	"context"

	"github.com/imirjar/Michman/config"
	"github.com/imirjar/Michman/internal/michman/gateway/http"
	"github.com/imirjar/Michman/internal/michman/service/grazer"
	"github.com/imirjar/Michman/internal/michman/service/reporter"
	"github.com/imirjar/Michman/internal/michman/storage/api"
	"github.com/imirjar/Michman/internal/michman/storage/memory"
)

func Run(ctx context.Context) error {
	// Подгружаем файл конфигураций /config
	cfg := config.New()

	// Приложение состоит из 3 основных слоев
	// 1)rpc - Шлюз через который приложений обрабатывает входящие запросы
	// 2)service - Сервис для логики обработки запросов
	// 3)storage - Хранилище значений
	// Слои представляют собой цепочку последовательностей
	apiStore := api.New()
	memStore := memory.New()

	grazer := grazer.New()
	reporter := reporter.New()

	srv := http.New()

	// Соединяем srv->service->storage
	// Так данные и будут двигаться
	grazer.Storage = memStore

	reporter.DiverStore = apiStore
	reporter.DiverStore = apiStore

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
