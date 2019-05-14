package main

import (
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/vladkampov/url-shortener-web-api/domain"
	"github.com/vladkampov/url-shortener-web-api/router"
	"net/http"
)

func main() {
	log.Printf("We are about to go...")
	domain.InitDomainGRPCSession()

	log.Println("Service has started at http://localhost:80")
	handler := cors.Default().Handler(router.InitRouter())
	log.Fatal(http.ListenAndServe(":80", handler))
}
