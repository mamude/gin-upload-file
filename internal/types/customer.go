package types

import (
	"fmt"
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
	} else {
		validCpf := fmt.Sprintf("%v", cpf)
		c.CPF = validCpf
	}
}

func (c *Customer) ValidateMostFrequentStore(value string) {
	cnpj := helpers.NewCNPJ(value)
	if !cnpj.IsValid() {
		c.MostFrequentStore = "invalid"
	} else {
		validCnpj := fmt.Sprintf("%v", cnpj)
		c.MostFrequentStore = validCnpj
	}
}

func (c *Customer) ValidateLastPurchaseStore(value string) {
	cnpj := helpers.NewCNPJ(value)
	if !cnpj.IsValid() {
		c.LastPurchaseStore = "invalid"
	} else {
		validCnpj := fmt.Sprintf("%v", cnpj)
		c.LastPurchaseStore = validCnpj
	}
}
