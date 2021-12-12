package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/richardsplit/go_translator_gopher/pkg/app"
	"github.com/richardsplit/go_translator_gopher/pkg/env"
	"github.com/richardsplit/go_translator_gopher/pkg/handlers"
	"github.com/richardsplit/go_translator_gopher/pkg/history"
	"github.com/richardsplit/go_translator_gopher/pkg/translation"

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

	sentenceHandler := handlers.NewSentenceHandler()
	router.HandleFunc("/sentence", handlers.TranslatorHandlerWrapperFunc(translator, sentenceHandler, history)).Methods(http.MethodPost)

	historyHandler := handlers.NewHistoryHandler()

	router.HandleFunc("/history", handlers.HandlerWrapperFunc(historyHandler, history)).Methods(http.MethodGet)

	portInfo := flag.Int("port", 8080, "port")
	flag.Parse()

	host := env.OptionalString("HOST", "localhost")

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

	logrus.Info("shutting down the server...")

	if err := httpServer.Shutdown(context.Background()); err != nil {
		logrus.WithError(err).Error("failed to properly stop http server")
	}
}
