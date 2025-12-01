package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"avito_project/course-go-avito-Vodyiara/internal/handler"
	"avito_project/course-go-avito-Vodyiara/internal/repository/postgres"
	"avito_project/course-go-avito-Vodyiara/internal/service"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

type Server struct {
	httpServer *http.Server
	db         *sql.DB
}

func New(port, dsn string) (*Server, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to DB successfully")

	courierRepo := postgres.NewCourierRepository(db)
	courierService := service.NewCourierService(courierRepo)
	courierHandler := handler.NewCourierHandler(courierService)

	r := chi.NewRouter()

	r.Get("/ping", handler.Ping)
	r.Head("/healthcheck", handler.Healthcheck)

	r.Get("/courier/{id}", courierHandler.GetCourier)
	r.Get("/couriers", courierHandler.GetAllCouriers)
	r.Post("/courier", courierHandler.CreateCourier)
	r.Put("/courier", courierHandler.UpdateCourier)

	addr := fmt.Sprintf(":%s", port)
	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: r,
		},
		db: db,
	}, nil
}

func (s *Server) Start() error {
	log.Printf("Server starting on %s...", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Shutting down server...")

	if err := s.db.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}

	return s.httpServer.Shutdown(ctx)
}
