package main

import (
	"context"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	initConfig()

	// get configuration from viper
	addr := viper.GetString("server.addr")

	// gin.New() unlike gin.Default(), (conceptually) does not attach any middleware to the Engine.
	r := gin.New()
	p := ginprom.New(ginprom.Engine(r)) // for prometheus
	r.Use(
		p.Instrument(),
		gin.Recovery(),
		RequestIDMiddleware(),
		// ginutil.LogrusLogger(),
		// ginutil.OtelMiddleware("backup-job-worker"),
		// ginutil.CorsMiddleware(),
		// ...
	)

	// add handlers
	// r = delivery.NewRouter(r)

	// ref: github.com/gin-gonic/examples/tree/master/graceful-shutdown/graceful-shutdown/notify-with-context
	// create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	svr := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	svr.BaseContext = func(net.Listener) context.Context {
		return ctx // for canceling running jobs
	}

	go func() {
		logrus.Println("Serving requests...")
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Failed to listen and serve: %v", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()
	logrus.Info("Shutting down server...")

	// this context is used to inform the server that it has 5 seconds to finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := svr.Shutdown(ctx); err != nil {
		logrus.Fatalf("Failed to shutdown server: %v", err)
	}
	logrus.Println("Server shutdown")
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.SetDefault("server.port", 8080)
	viper.SetDefault("storage.dir", "./data")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Info("no config file found; using default settings")
		} else {
			logrus.Panicf("failed to read config file: %v", err.Error())
		}
	}
	viper.AutomaticEnv()
}
