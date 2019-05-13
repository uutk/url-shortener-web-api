package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/vladkampov/url-shortener-web-api/domain"
	"github.com/vladkampov/url-shortener-web-api/router"
	"net/http"
)

func main() {
	log.Printf("We are about to go...")
	domain.InitDomainGRPCSession()

	log.Println("Service has started at http://localhost:80")
	log.Fatal(http.ListenAndServe(":80", router.InitRouter()))
}
