package handler

import (
	"context"
	"fmt"
	"friend-management-go/internal/config"
	"friend-management-go/internal/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
)

type Server struct {
	cfg    config.HTTPServer
	repo   repository.IFriendshipRepository
	router *chi.Mux
}

func NewServer(cfg config.HTTPServer, repo repository.IFriendshipRepository) *Server {
	srv := &Server{
		cfg:    cfg,
		repo:   repo,
		router: chi.NewRouter(),
	}

	srv.routes()

	return srv
}

func (s *Server) Start(c context.Context) {
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", s.cfg.Port),
		Handler:      s.router,
		IdleTimeout:  s.cfg.IdleTimeout,
		ReadTimeout:  s.cfg.ReadTimeout,
		WriteTimeout: s.cfg.WriteTimeout,
	}

	shutdownComplete := handleShutdown(func() {
		if err := server.Shutdown(c); err != nil {
			log.Printf("server.Shutdown failed: %v\n", err)
		}
	})

	if err := server.ListenAndServe(); err == http.ErrServerClosed {
		<-shutdownComplete
	} else {
		log.Printf("http.ListenAndServe failed: %v\n", err)
	}

	log.Println("Shutdown gracefully")
}

func handleShutdown(onShutdownSignal func()) <-chan struct{} {
	shutdown := make(chan struct{})

	go func() {
		shutdownSignal := make(chan os.Signal, 1)
		signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

		<-shutdownSignal

		onShutdownSignal()
		close(shutdown)
	}()

	return shutdown
}
