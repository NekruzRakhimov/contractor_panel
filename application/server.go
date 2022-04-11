package application

import (
	"context"
	"contractor_panel/application/config"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/spf13/viper"
)

// Server provides an http.Server.
type Server struct {
	*http.Server
}

// NewServer creates and configures an APIServer serving all application routes.
func NewServer() (*Server, error) {
	err := config.CheckEnv()
	if err != nil {
		return nil, err
	}
	api, err := NewApi()
	if err != nil {
		return nil, err
	}

	var addr string
	port := viper.GetString(config.Port)

	if strings.Contains(port, ":") {
		addr = port
	} else {
		addr = ":" + port
	}

	srv := http.Server{
		Addr:    addr,
		Handler: api,
	}

	return &Server{&srv}, nil
}

// Start производит запуск сервиса на указанном порту
func (srv *Server) Start() {
	log.Info("starting server...")
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Infof("Listening on %s", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Info("Shutting down server... Reason:", sig)
	// teardown logic...

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Info("Server gracefully stopped")
}
