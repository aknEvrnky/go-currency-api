package today

import (
	"encoding/json"
	"github.com/aknevrnky/go-currency-api/pkg/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type TodayApi struct {
	Service *service.CurrencyService
}

func NewTodayApi(service service.CurrencyService) *TodayApi {
	return &TodayApi{Service: &service}
}

func (api *TodayApi) GetAllToday(resp http.ResponseWriter, req *http.Request) {
	codes, err := (*api.Service).GetAll()
	if err != nil {
		log.Printf("Error getting currencies: %v", err)

		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Error getting currencies"))
		return
	}

	resp.Header().Set("Content-Type", "application/json")

	// convert map to json
	data, err := json.Marshal(codes)

	if err != nil {
		log.Printf("Error marshalling currencies: %v", err)

		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Error marshalling currencies"))
		return
	}

	resp.Write(data)
	return
}

func (api *TodayApi) GetByCodeToday(resp http.ResponseWriter, req *http.Request) {
	code := mux.Vars(req)["code"]

	if code == "" {
		log.Printf("Error getting currencies: %v", "code is required")

		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("code is required"))
		return
	}

	currency, err := (*api.Service).GetByCode(code)
	if err != nil {
		log.Printf("Error getting currencies: %v", err)

		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Error getting currencies"))
		return
	}

	resp.Header().Set("Content-Type", "application/json")

	// convert map to json
	data, err := json.Marshal(currency)

	if err != nil {
		log.Printf("Error marshalling currencies: %v", err)

		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Error marshalling currencies"))
		return
	}

	resp.Write(data)
	return
}
