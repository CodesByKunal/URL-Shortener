package main

import (
	"log"
	"net/http"
	"os"
	"urlshortener/handler"
	"urlshortener/middleware"
)

var (
	HOST string = os.Getenv("HOST")
	PORT string = os.Getenv("PORT")
)

func main() {

	if HOST == "" || PORT == "" {
		log.Fatal("Environment variables HOST and PORT must be set")
	}

	addr := HOST + ":" + PORT

	http.HandleFunc("/health", handler.HealthHandler)

	err := http.ListenAndServe(addr, middleware.Logs(http.DefaultServeMux))
	if err != nil {
		log.Fatal("Unable to Start Server", err)
	}

}
