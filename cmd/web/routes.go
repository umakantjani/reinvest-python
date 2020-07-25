package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func (app *application) routes() http.Handler {
	// Create a new router
	router := mux.NewRouter()

	router.HandleFunc("/", app.welcome).Methods("GET")
	//router.HandleFunc("/collect/kpis/{kpi}/{symbol}/{period}", app.collectProfile).Methods("POST")
	router.HandleFunc("/collect/kpis/{kpi}/{symbol}", app.collectKpiData).Methods("POST")
	router.HandleFunc("/collect/kpis/{kpi}/{symbol}/{period}", app.collectKpiData).Methods("POST")
	router.HandleFunc("/collect/lists/{listName}", app.collectLists).Methods("POST")
	router.HandleFunc("/find", app.findAllDoc).Methods("GET")
	router.HandleFunc("/find/{code}", app.findSingleDoc).Methods("GET")

	handler := cors.Default().Handler(router)

	return handler

}