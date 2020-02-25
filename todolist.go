package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// API Health Check
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Info("API Health Check Passed")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	log.Info("Starting TodoList API Server")
	router := mux.NewRouter()
	router.HandleFunc("/health_check", HealthCheck).Methods("GET")
	http.ListenAndServe(":8964", router)
}
