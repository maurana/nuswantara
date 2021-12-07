package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

  	"github.com/maurana/nuswantara/core/config"
	"github.com/maurana/nuswantara/core/routes"
	"github.com/maurana/nuswantara/core/log"
)

func Startup() error {
	httpServer := &http.Server{
	  Addr:    fmt.Sprintf(":%d", config.Cfg().AppPort),
	  Handler: NuswantaraRouter(),
	}
  
	idleConnsClosed := make(chan struct{})
	  go func() {
		  defer close(idleConnsClosed)
  
		  sigint := make(chan os.Signal, 1)
		  signal.Notify(sigint, os.Interrupt)
		  signal.Notify(sigint, syscall.SIGTERM)
  
		  <-sigint
  
		  err := httpServer.Shutdown(context.Background())
		  if err != nil {
			  log.Log().Err(err).Msg("failed to shutdown server")
		  }
	  }()
  
	  log.Log().Info().Msgf("starting server on port%s", httpServer.Addr)
	  err := httpServer.ListenAndServe()
	  if err != nil && err != http.ErrServerClosed {
		  return err
	  }
  
	  <-idleConnsClosed
  
	  log.Log().Info().Msg("stopped server gracefully")
	  return nil
  }