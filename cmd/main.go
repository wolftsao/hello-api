package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/wolftsao/hello-api/handlers"
	"github.com/wolftsao/hello-api/handlers/rest"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080"
	}

	mux := http.NewServeMux()

	mux.Handle("/translate/hello", http.StripPrefix("/translate", http.HandlerFunc(rest.TranslateHandler)))
	mux.HandleFunc("/health", handlers.HealthCheck)

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}
