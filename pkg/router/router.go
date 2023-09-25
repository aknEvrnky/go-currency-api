package router

import (
	"github.com/aknevrnky/go-currency-api/pkg/api/today"
	"github.com/aknevrnky/go-currency-api/pkg/repository"
	"github.com/aknevrnky/go-currency-api/pkg/service"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func New(rdb *redis.Client) *mux.Router {
	router := mux.NewRouter()

	tcmbRepo := repository.NewTcmbRepository(rdb)
	tcmbService := service.NewTcmbService(tcmbRepo)
	todayApi := today.NewTodayApi(tcmbService)

	setupRoutes(router, todayApi)

	return router
}

func setupRoutes(router *mux.Router, todayApi *today.TodayApi) {
	router.HandleFunc("/today", todayApi.GetAllToday).Methods("GET")
	router.HandleFunc("/today/{code}", todayApi.GetByCodeToday).Methods("GET")
}
