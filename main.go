package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/richardsplit/translator_go/pkg/app"
	"github.com/richardsplit/translator_go/pkg/env"
	"github.com/richardsplit/translator_go/pkg/handlers"
	"github.com/richardsplit/translator_go/pkg/history"
	"github.com/richardsplit/translator_go/pkg/translation"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host string
	Port int
}

func main() {
	lifecycle := app.NewLifecycle()

	translator := translation.NewTranslator()

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.HomeHandler).Methods(http.MethodGet)

	history := history.History()

	wordHandler := handlers.NewWordHandler()

	router.HandleFunc("/word", handlers.TranslatorHandlerWrapperFunc(translator, wordHandler, history)).
		Methods(http.MethodPost)

	historyHandler := handlers.NewHistoryHandler()

	router.HandleFunc("/history", handlers.HandlerWrapperFunc(historyHandler, history)).Methods(http.MethodGet)

	fmt.Printf("Enter the port for the application: ")
	var first int
	fmt.Scanln(&first)
	portInfo := flag.Int("port", first, "port")
	flag.Parse()

	host := env.OptionalString("HOST", "127.0.0.1")

	port, err := env.OptionalInt("PORT", *portInfo)

	if err != nil {
		lifecycle.Crash(errors.Wrap(err, "failed to load config"))
	}

	config := Config{
		Host: host,
		Port: port,
	}

	logrus.Info("starting http server...")
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler: router,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			lifecycle.Crash(errors.Wrap(err, "http server returned an error"))
			logrus.WithError(err).Fatal("application has crashed")

		}
	}()

	lifecycle.WaitExitSignal()

	logrus.Info("shutting down the server,please wait...")

	if err := httpServer.Shutdown(context.Background()); err != nil {
		logrus.WithError(err).Error("failed to properly stop http server")
	}
}
