package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/abdil1234/test-golang/internal/app/commons/loggers"

	"github.com/abdil1234/test-golang/internal/app/commons"
	"github.com/abdil1234/test-golang/internal/app/controllers"
	"github.com/abdil1234/test-golang/internal/app/usecases"
)

var logger = loggers.GetLogger(context.Background())

type IServer interface {
	StartApp()
}

type server struct {
	opt      commons.Option
	services *usecases.Services
}

func NewServer(opt commons.Option, services *usecases.Services) IServer {
	return &server{
		opt:      opt,
		services: services,
	}
}

func (s *server) StartApp() {
	var srv http.Server
	idleConnectionClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		logger.Infoln("[API] Server is shutting down")

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			logger.Errorf("[API] Fail to shutting down: %v", err)
		}
		close(idleConnectionClosed)
	}()

	srv.Addr = fmt.Sprintf("%s:%d", s.opt.AppCtx.GetAppOption().Host, s.opt.AppCtx.GetAppOption().Port)
	hOpt := controllers.ControllerOption{
		Option:   s.opt,
		Services: s.services,
	}
	srv.Handler = Router(hOpt)

	logger.Infof("[API] HTTP serve at %s\n", srv.Addr)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		logger.Fatalf("[API] Fail to start listen and server: %v", err)
	}

	<-idleConnectionClosed
	logger.Infoln("[API] Bye")
}
