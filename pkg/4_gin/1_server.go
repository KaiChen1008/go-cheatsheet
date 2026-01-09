package ginutil

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// gin.New() unlike gin.Default(), (conceptually) does not attach any middleware to the Engine.
	r := gin.New()

	p := ginprom.New(ginprom.Engine(r)) // for prometheus
	r.Use(
		p.Instrument(),
		gin.Recovery(),
		RequestIDMiddleware(),
	)

	// add handlers
	// r = delivery.NewRouter(r)

	svr := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// graceful shutdown
	go func() {
		logrus.Println("Serving requests...")
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Failed to listen and serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM) // SIGINT & SIGTERM
	<-quit

	logrus.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := svr.Shutdown(ctx); err != nil {
		logrus.Fatalf("Failed to shutdown server: %v", err)
	}
	logrus.Println("Server shutdown")

}
