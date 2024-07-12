package service

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"example.com/mamude/internal/helpers"
	"example.com/mamude/internal/types"
)

const (
	cpf                = 0
	private            = 1
	incomplete         = 2
	dateLastPurchase   = 3
	averageTicket      = 4
	lastPurchaseTicket = 5
	mostFrequentStore  = 6
	lastPurchaseStore  = 7
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("unable to read file %v", err)
	}
}

func SanitizeData(fileName string) []types.Customer {
	file, err := os.Open(fileName)
	checkError(err)

	// fechar o arquivo
	defer file.Close()

	// iniciar a leitura do arquivo
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	// ignorar a primeira linha (headers)
	scanner.Scan()
	customers := []types.Customer{}

	for scanner.Scan() {
		field := strings.Fields(scanner.Text())

		// parsers
		cpf := helpers.ParseCPF(field[cpf])
		private, _ := strconv.Atoi(field[private])
		incomplete, _ := strconv.Atoi(field[incomplete])
		dateLastPurchase := helpers.ParseDate(field[dateLastPurchase])
		averageTicket := helpers.ParseDecimal(field[averageTicket])
		lastPurchaseTicket := helpers.ParseDecimal(field[lastPurchaseTicket])
		mostFrequentStore := helpers.ParseCNPJ(field[mostFrequentStore])
		lastPurchaseStore := helpers.ParseCNPJ(field[lastPurchaseStore])

		customers = append(customers,
			types.Customer{
				CPF:                cpf,
				Private:            private,
				Incomplete:         incomplete,
				DateLastPurchase:   dateLastPurchase,
				AverageTicket:      averageTicket,
				LastPurchaseTicket: lastPurchaseTicket,
				MostFrequentStore:  mostFrequentStore,
				LastPurchaseStore:  lastPurchaseStore,
			})
	}

	return customers
}
