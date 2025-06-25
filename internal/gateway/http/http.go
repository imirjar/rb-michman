package http

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/imirjar/rb-michman/internal/models"
)

type Service interface {
	GetReports(context.Context) ([]models.Report, error)
	ExecuteReport(context.Context, string) (models.Data, error)
	// ExecuteReportMap(ctx context.Context, hash string) (string, error)
}

type App struct {
	Service Service
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
	router.Get("/info", a.Info())

	// Get available divers
	router.Route("/reports", func(reports chi.Router) {
		reports.Get("/", a.GetReports()) // diver reports list from DB
		reports.Post("/{id}", a.ExecuteReport())
	})

	// router.Post("/connect", a.ConnectDiver())

	srv := &http.Server{
		Addr:    ":" + addr,
		Handler: router,
	}

	log.Printf("Start app on PORT %s", addr)
	return srv.ListenAndServe()
}
