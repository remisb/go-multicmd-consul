package main

import (
	"fmt"
	"go-multicmd-docker/internal/config"
	"go-multicmd-docker/internal/health"
	"log"
	"net/http"
)

const APP_NAME = "Server"

func main() {
	cfg := config.NewFromEnv()

	mux := http.NewServeMux()
	mux.HandleFunc("/server", indexHander)
	mux.HandleFunc("/server/health", health.HealthHandler)

	fmt.Println(APP_NAME, "started and accessible at localhost:", cfg.ServerAddress)

	err := http.ListenAndServe(cfg.ServerAddress, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func indexHander(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("server index page"))
}
