package service

import "github.com/aknevrnky/go-currency-api/pkg/entity/currency"

type CurrencyService interface {
	GetAll() (map[string]currency.Currency, error)
	GetByCode(code string) (*currency.Currency, error)
}
