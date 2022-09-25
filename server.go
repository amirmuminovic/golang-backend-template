package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func startServer(ac AppConfig) {
	http.HandleFunc("/health", handleHealthCheck)
	log.Fatal(http.ListenAndServe(":"+ac.appPort, nil))
}
