package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imirjar/Michman/config"
	"github.com/imirjar/Michman/internal/michman/models"
	"github.com/imirjar/Michman/internal/michman/service"
	"github.com/imirjar/rb-glue/middlewares/authentication"
	"github.com/imirjar/rb-glue/middlewares/contype"
	"github.com/imirjar/rb-glue/middlewares/logger"
)

type Service interface {
	DiverList(context.Context) ([]models.Diver, error)
	DiverReports(context.Context, string) ([]models.Report, error)
	GetDiverReportData(ctx context.Context, ex models.Diver) (models.Report, error)
}

type App struct {
	config  Config
	Service Service
}

type Config interface {
	GetAuthAddr() string
	GetMichmanAddr() string //allow req only for this addr
	GetSecret() string
}

func New() *App {

	return &App{
		config:  config.NewConfig(),
		Service: service.New(),
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

	router.Post("/divers/", a.DiversList())

	router.Route("/diver", func(diver chi.Router) {
		diver.Post("/reports/", a.ReportsList())
		diver.Post("/execute/", a.ReportExecute())
	})

	router.Route("/connect", func(conn chi.Router) {
		conn.Post("/diver/", a.ReportsList())
	})

	srv := &http.Server{
		Addr:    a.config.GetMichmanAddr(),
		Handler: router,
	}

	log.Printf("Run app on PORT %s", a.config.GetMichmanAddr())
	return srv.ListenAndServe()
}
