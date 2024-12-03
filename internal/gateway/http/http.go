package http

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/imirjar/rb-michman/internal/models"
)

type Diver interface {
	DiverReports(context.Context, string) ([]models.Report, error)
	GetDiverReportData(ctx context.Context, addr, repID string) (models.Data, error)
	GetDiverReportDataMap(ctx context.Context, addr, repID string) ([]map[string]interface{}, error)
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
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

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
