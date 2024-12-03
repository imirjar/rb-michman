package http

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imirjar/rb-michman/internal/models"
)

type Diver interface {
	DiverReports(context.Context, string) ([]models.Report, error)
	GetDiverReportData(ctx context.Context, addr, repID string) (models.Data, error)
}

type Grazer interface {
	ConnectDiver(context.Context, models.Diver) error
	CheckConnections(context.Context) error // read all connected divers, ping it, connect which is still alive
	DiverList(context.Context) (map[string]models.Diver, error)
	DiverAddr(ctx context.Context, hash string) (string, error)
}

type App struct {
	ReportService Diver
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
	router.Route("/divers", func(diver chi.Router) {
		diver.Get("/", a.GetDivers())
		diver.Get("/{id}", a.GetDiverReports())
		diver.Get("/{id}/{reportId}", a.ExecuteDiverReport())
	})

	router.Post("/connect", a.ConnectDiver())

	log.Println(addr)
	srv := &http.Server{
		Addr:    ":" + addr,
		Handler: router,
	}

	log.Printf("Start app on PORT %s", addr)
	return srv.ListenAndServe()
}
