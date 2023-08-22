package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/wolftsao/hello-api/handlers"
	"github.com/wolftsao/hello-api/handlers/rest"
	"github.com/wolftsao/hello-api/translation"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080"
	}

	mux := http.NewServeMux()

	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)

	mux.Handle("/translate/hello", http.StripPrefix("/translate", http.HandlerFunc(translateHandler.TranslateHandler)))
	mux.HandleFunc("/health", handlers.HealthCheck)

	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           mux,
	}

	log.Printf("listening on %s\n", addr)

	log.Fatal(server.ListenAndServe())
}
