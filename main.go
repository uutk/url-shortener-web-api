package main

import (
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/vladkampov/url-shortener-web-api/domain"
	"github.com/vladkampov/url-shortener-web-api/router"
	"net/http"
	"github.com/joho/godotenv"
)

func main() {
	log.Printf("We are about to go...")
	err := godotenv.Load()
	if err != nil {
		log.Warn("Error loading .env file")
	}
	domain.InitDomainGRPCSession()

	log.Println("Service has started at http://localhost:80")
	handler := cors.Default().Handler(router.InitRouter())
	log.Fatal(http.ListenAndServe(":80", handler))
}
