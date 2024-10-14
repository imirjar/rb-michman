package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imirjar/Michman/config"
	"github.com/imirjar/Michman/internal/michman/models"
	"github.com/imirjar/Michman/internal/michman/service/grazer"
	"github.com/imirjar/Michman/internal/michman/service/reporter"
	"github.com/imirjar/rb-glue/middlewares/authentication"
	"github.com/imirjar/rb-glue/middlewares/contype"
	"github.com/imirjar/rb-glue/middlewares/logger"
)

type Reporter interface {
	DiverReports(context.Context, string) ([]models.Report, error)
	GetDiverReportData(ctx context.Context, ex models.Diver) (models.Report, error)
}

type Grazer interface {
	LoadConnections()   // read all connected divers, ping it, connect which is still alive
	BackupConnections() // backup all connected divers connection info
	DiverList(context.Context) ([]models.Diver, error)
}

type App struct {
	config        Config
	ReportService Reporter
	GrazerService Grazer
}

type Config interface {
	GetAuthAddr() string
	GetMichmanAddr() string //allow req only for this addr
	GetSecret() string
}

func New() *App {

	return &App{
		config:        config.NewConfig(),
		ReportService: reporter.New(),
		GrazerService: grazer.New(),
	}
}

func (a *App) Run(ctx context.Context) error {
	router := chi.NewRouter()

	authpath := fmt.Sprintf(a.config.GetAuthAddr() + "/token/validate")

	// Middlewares
	router.Use(authentication.Authenticate(authpath, authentication.UserParams{}))
	router.Use(logger.Logger())
	router.Use(contype.CheckType("application/json"))

	// Check connection
	router.Get("/", a.Ping())

	// Get available divers
	router.Get("/divers", a.DiversList())

	router.Route("/diver/{id}", func(diver chi.Router) {
		diver.Get("/", a.ReportsList())
		diver.Post("/execute/{reportId}", a.ReportExecute())
	})

	router.Post("/connect", a.ReportsList())

	srv := &http.Server{
		Addr:    a.config.GetMichmanAddr(),
		Handler: router,
	}

	log.Printf("Run app on PORT %s", a.config.GetMichmanAddr())
	return srv.ListenAndServe()
}
