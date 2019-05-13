package router

import (
	//"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vladkampov/url-shortener-web-api/domain"

	//"github.com/vladkampov/url-shortener-web-api/domain"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func redirectToUrl(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	url, err := domain.GetUrl(vars["hash"])

	if err != nil {
		log.Warn(err)
		http.Redirect(w, r, "http://kampov.com", http.StatusMovedPermanently)
		return
	}

	if len(url) != 0 {
		log.Printf("Getting to %s", url)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	}
}

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/{hash}", redirectToUrl).Methods("GET")

	return router
}
