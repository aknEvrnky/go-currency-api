package repository

import "github.com/aknevrnky/go-currency-api/pkg/entity/currency"

type CurrencyRepository interface {
	GetAll() (map[string]currency.Currency, error)
	GetByCode(code string) (*currency.Currency, error)
}
