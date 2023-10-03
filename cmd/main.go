package main

import (
	"log"
	"net/http"
	"time"

	"github.com/wolftsao/hello-api/config"
	"github.com/wolftsao/hello-api/handlers"
	"github.com/wolftsao/hello-api/handlers/rest"
	"github.com/wolftsao/hello-api/translation"
)

func main() {
	cfg := config.LoadConfiguration()
	addr := cfg.Port

	mux := API(cfg)

	log.Printf("listening on %s\n", addr)

	srv := http.Server{
		ReadTimeout: time.Second * 10,
		Handler:     mux,
		Addr:        addr,
	}

	// log.Fatal(http.ListenAndServe(addr, mux))
	log.Fatal(srv.ListenAndServe())
}

func API(cfg config.Configuration) *http.ServeMux {
	mux := http.NewServeMux()

	var translationService rest.Translator
	translationService = translation.NewStaticService()
	if cfg.LegacyEndpoint != "" {
		log.Printf("creating external translation client: %s", cfg.LegacyEndpoint)
		client := translation.NewHelloClient(cfg.LegacyEndpoint)
		translationService = translation.NewRemoteService(client)
	}

	if cfg.DatabaseURL != "" {
		db := translation.NewDatabaseService(cfg)
		translationService = db
	}

	translateHandler := rest.NewTranslateHandler(translationService)

	mux.Handle("/translate/hello", http.StripPrefix("/translate", http.HandlerFunc(translateHandler.TranslateHandler)))
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/info", handlers.Info)

	return mux
}
