package main

import (
	"fmt"
	"go-multicmd-docker/internal/config"
	"go-multicmd-docker/internal/health"
	"log"
	"net/http"
)

const APP_NAME = "Admin"

func main() {
	cfg := config.NewFromEnv()

	mux := http.NewServeMux()
	mux.HandleFunc("/admin", indexHander)
	mux.HandleFunc("/admin/health", health.HealthHandler)

	fmt.Println(APP_NAME, "started and accessible at localhost:", cfg.AdminAddress)

	err := http.ListenAndServe(cfg.AdminAddress, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func indexHander(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("admin index page"))
}
