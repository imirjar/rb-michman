package http

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imirjar/Michman/internal/michman/models"
)

type Reporter interface {
	DiverReports(context.Context, string) ([]models.Report, error)
	GetDiverReportData(ctx context.Context, ex models.Diver) (models.Report, error)
}

type Grazer interface {
	ConnectDiver(context.Context, models.Diver) error
	LoadConnections()   // read all connected divers, ping it, connect which is still alive
	BackupConnections() // backup all connected divers connection info
	DiverList(context.Context) ([]models.Diver, error)
}

type App struct {
	ReportService Reporter
	GrazerService Grazer
}

func New() *App {

	return &App{}
}

func (a *App) Start(ctx context.Context, addr string) error {
	router := chi.NewRouter()

	// authpath := fmt.Sprintf(a.config.GetAuthAddr() + "/token/validate")

	// Middlewares
	// router.Use(authentication.Authenticate(authpath, authentication.UserParams{}))
	// router.Use(logger.Logger())
	// router.Use(contype.CheckType("application/json"))

	// Check connection
	router.Get("/", a.Info())

	// Get available divers
	router.Get("/divers", a.Grazer())

	router.Route("/diver/{id}", func(diver chi.Router) {
		diver.Get("/", a.ReportsList())
		diver.Post("/execute/{reportId}", a.ReportExecute())
	})

	router.Post("/connect", a.Connect())

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Printf("Start app on PORT %s", addr)
	return srv.ListenAndServe()
}
