package repository

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/xml"
	"github.com/aknevrnky/go-currency-api/pkg/entity/currency"
	"github.com/redis/go-redis/v9"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	apiEndpoint = "https://www.tcmb.gov.tr/kurlar/today.xml"
)

var (
	ctx = context.Background()
)

type TcmbRepository struct {
	RDB *redis.Client
}

type apiResponse struct {
	Currencies []currency.Currency `xml:"Currency"`
}

func NewTcmbRepository(rdb *redis.Client) *TcmbRepository {
	return &TcmbRepository{RDB: rdb}
}

func (t *TcmbRepository) GetAll() (map[string]currency.Currency, error) {
	var apiRes apiResponse

	// check if currencies are cached
	byteData, err := t.RDB.Get(ctx, "currencies").Bytes()
	if err == redis.Nil {
		log.Println("Cache does not exist. Fetching from API...")

		// fetch from api
		res, err := fetchApi()
		apiRes = *res
		if err != nil {
			return nil, err
		}

		// convert data to bytes
		var buffer bytes.Buffer
		encoder := gob.NewEncoder(&buffer)
		err = encoder.Encode(apiRes)
		if err != nil {
			return nil, err
		}
		byteData := buffer.Bytes()

		// save to cache
		err = t.RDB.Set(ctx, "currencies", byteData, time.Minute).Err()

	} else if err != nil {
		return nil, err
	} else {
		log.Println("Cache exists. Fetching from cache...")

		// convert bytes to data
		buffer := bytes.NewBuffer(byteData)
		decoder := gob.NewDecoder(buffer)
		err = decoder.Decode(&apiRes)

		if err != nil {
			// delete cache, fetch from api and save to cache
			t.RDB.Del(ctx, "currencies")
			return t.GetAll()
		}
	}

	// convert currencies to map
	currencies := make(map[string]currency.Currency, len(apiRes.Currencies))
	for _, c := range apiRes.Currencies {
		currencies[c.Code] = c
	}

	// return response
	return currencies, nil
}

func (t *TcmbRepository) GetByCode(code string) (*currency.Currency, error) {
	// get all currencies
	currencies, err := t.GetAll()
	if err != nil {
		return nil, err
	}

	// get currency by code
	c, ok := currencies[code]
	if !ok {
		return nil, nil
	}

	// return response
	return &c, nil
}

func fetchApi() (*apiResponse, error) {
	var apiRes apiResponse

	// make a get request to apiEndpoint
	resp, err := http.Get(apiEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parse xml response
	err = xml.Unmarshal(data, &apiRes)
	if err != nil {
		return nil, err
	}

	return &apiRes, nil
}
