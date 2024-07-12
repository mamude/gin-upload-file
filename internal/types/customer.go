package types

import (
	"time"

	"example.com/mamude/internal/helpers"
)

type Customer struct {
	CPF                string
	Private            int
	Incomplete         int
	DateLastPurchase   time.Time
	AverageTicket      float64
	LastPurchaseTicket float64
	MostFrequentStore  string
	LastPurchaseStore  string
}

func (c *Customer) ValidateCPF(value string) {
	cpf := helpers.NewCPF(value)
	if !cpf.IsValid() {
		c.CPF = "invalid"
	}
	c.CPF = cpf.String()
}

func (c *Customer) ValidateMostFrequentStore(value string) {
	cnpj := helpers.NewCNPJ(value)
	if !cnpj.IsValid() {
		c.MostFrequentStore = "invalid"
	}
	c.MostFrequentStore = cnpj.String()
}

func (c *Customer) ValidateLastPurchaseStore(value string) {
	cnpj := helpers.NewCNPJ(value)
	if !cnpj.IsValid() {
		c.LastPurchaseStore = "invalid"
	}
	c.LastPurchaseStore = cnpj.String()
}
