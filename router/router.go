package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/vladkampov/url-shortener-web-api/domain"
	"net/http"
)

type UrlObject struct {
	Url string `json:"url"`
	UserId string `json:"userId"`
}

type ErrorResponse struct {
	Error string
}


func redirectToUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["hash"] == "favicon.ico" {
		return
	}

	url, err := domain.GetUrl(vars["hash"])

	if err != nil {
		log.Warn(err)
		http.Redirect(w, r, "http://kampov.com/url-shortener", http.StatusMovedPermanently)
		return
	}

	if len(url) != 0 {
		log.Printf("Getting to %s", url)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	}
}

func handleShortRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body UrlObject
	_ = json.NewDecoder(r.Body).Decode(&body)

	url, err := domain.SendUrl(body.Url, body.UserId)
	if err != nil {
		log.Warnf("Can't short the url %s: %s", body.Url, err)
		w.WriteHeader(500)
		err = json.NewEncoder(w).Encode(ErrorResponse{err.Error()})
		if err != nil {
			log.Warn(err)
		}
		return
	}

	err = json.NewEncoder(w).Encode(UrlObject{url, body.UserId})
	if err != nil {
		log.Warn(err)
	}
}

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/short-it", handleShortRequest).Methods("POST")
	router.HandleFunc("/{hash}", redirectToUrl).Methods("GET")

	return router
}
