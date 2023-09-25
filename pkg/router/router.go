package router

import (
	"github.com/aknevrnky/go-currency-api/pkg/api/today"
	"github.com/gorilla/mux"
)

func New(todayApi *today.TodayApi) *mux.Router {
	router := mux.NewRouter()

	setupRoutes(router, todayApi)

	return router
}

func setupRoutes(router *mux.Router, todayApi *today.TodayApi) {
	router.HandleFunc("/today", todayApi.GetAllToday).Methods("GET")
	router.HandleFunc("/today/{code}", todayApi.GetByCodeToday).Methods("GET")
}
