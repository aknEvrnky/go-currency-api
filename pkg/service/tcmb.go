package service

import (
	"github.com/aknevrnky/go-currency-api/pkg/entity/currency"
	"github.com/aknevrnky/go-currency-api/pkg/repository"
	"strings"
)

type TcmbService struct {
	Repository *repository.CurrencyRepository
}

func NewTcmbService(repository repository.CurrencyRepository) *TcmbService {
	return &TcmbService{Repository: &repository}
}

func (t *TcmbService) GetAll() (map[string]currency.Currency, error) {
	return (*t.Repository).GetAll()
}

func (t *TcmbService) GetByCode(code string) (*currency.Currency, error) {
	code = strings.ToUpper(code)
	return (*t.Repository).GetByCode(code)
}
