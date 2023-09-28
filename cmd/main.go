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

	mux := http.NewServeMux()

	var translationService rest.Translator
	translationService = translation.NewStaticService()
	if cfg.LegacyEndpoint != "" {
		log.Printf("creating external translation client: %s", cfg.LegacyEndpoint)
		client := translation.NewHelloClient(cfg.LegacyEndpoint)
		translationService = translation.NewRemoteService(client)
	}
	translateHandler := rest.NewTranslateHandler(translationService)

	mux.Handle("/translate/hello", http.StripPrefix("/translate", http.HandlerFunc(translateHandler.TranslateHandler)))
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/info", handlers.Info)

	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           mux,
	}

	log.Printf("listening on %s\n", addr)

	log.Fatal(server.ListenAndServe())
}
