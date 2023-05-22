package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
)

var homeTemplate *template.Template

type payload struct {
	AppVersion    string
	ExoplanetList []Exoplanet
}

func init() {
	homeTemplate = template.Must(template.ParseFiles("template.html"))
}

func listenAndServe(port string, exoplanets *Exoplanets) error {

	var (
		template_payload payload
	)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		exoplanets.fetch()

		template_payload.AppVersion = version()
		template_payload.ExoplanetList = exoplanets.List

		w.WriteHeader(http.StatusOK)
		err := homeTemplate.Execute(w, template_payload)
		if err != nil {
			log.Fatalf("execution failed: %s", err)
			panic(err)
		}
	})

	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Printf("Listening on :%s", port)
	return srv.ListenAndServe()
}
