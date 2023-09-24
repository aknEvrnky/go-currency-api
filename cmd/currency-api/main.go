package main

import (
	"context"
	"github.com/aknevrnky/go-currency-api/pkg/api/today"
	"github.com/aknevrnky/go-currency-api/pkg/repository"
	"github.com/aknevrnky/go-currency-api/pkg/service"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
)

type Application struct {
	RDB    *redis.Client
	Router *mux.Router
}

func main() {
	app := bootstrap()

	app.AssignRoutes()

	app.Run(":8080")
}

func bootstrap() Application {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// create redis client
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
		Protocol: 3, // specify 2 for RESP 2 or 3 for RESP 3
	})

	// Ping Redis
	_, err = rdb.Ping(ctx).Result()

	if err != nil {
		log.Fatal("Error connecting to Redis")
	}

	// create router
	router := mux.NewRouter()

	// Create app
	app := Application{
		RDB:    rdb,
		Router: router,
	}

	return app
}

func (a *Application) AssignRoutes() {
	tcmbRepo := repository.NewTcmbRepository(a.RDB)
	tcmbService := service.NewTcmbService(tcmbRepo)
	todayApi := today.NewTodayApi(tcmbService)

	a.Router.HandleFunc("/today", todayApi.GetAllToday).Methods("GET")
	a.Router.HandleFunc("/today/{code}", todayApi.GetByCodeToday).Methods("GET")
}

func (a *Application) Run(port string) {
	log.Printf("Listening the server on port: %s\n", port)
	err := http.ListenAndServe(port, a.Router)
	if err != nil {
		log.Fatalf("Error starting server %s\n", err.Error())
	}
}
