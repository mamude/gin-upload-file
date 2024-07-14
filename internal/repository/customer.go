package repository

import (
	"example.com/mamude/internal/database"
	"example.com/mamude/internal/helpers"
	"example.com/mamude/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type CustomerRepository struct {
	db database.PgxIface
}

func NewCustomerRepository(db database.PgxIface) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (c *CustomerRepository) SaveData(ctx *gin.Context, customers []types.Customer) int64 {
	// tratar os dados e validá-los
	if len(customers) > 0 {
		var customersValid []types.Customer
		for _, customer := range customers {
			changedCustomer := &customer
			changedCustomer.ValidateCPF(customer.CPF)
			changedCustomer.ValidateMostFrequentStore(customer.MostFrequentStore)
			changedCustomer.ValidateLastPurchaseStore(customer.LastPurchaseStore)
			customersValid = append(customersValid, customer)
		}

		// enviar os dados em batch, otimizando o desempenho
		// https://pkg.go.dev/github.com/jackc/pgx/v5@v5.3.0#hdr-Copy_Protocol
		copyCount, err := c.db.CopyFrom(
			ctx,
			pgx.Identifier{"customers"},
			[]string{
				"cpf",
				"private",
				"incomplete",
				"date_last_purchase",
				"average_ticket",
				"last_purchase_ticket",
				"most_frequent_store",
				"last_purchase_store",
			},
			pgx.CopyFromSlice(len(customers), func(i int) ([]any, error) {
				return []any{
					customersValid[i].CPF,
					customersValid[i].Private,
					customersValid[i].Incomplete,
					customersValid[i].DateLastPurchase,
					customersValid[i].AverageTicket,
					customersValid[i].LastPurchaseTicket,
					customersValid[i].MostFrequentStore,
					customersValid[i].LastPurchaseStore,
				}, nil
			}),
		)
		helpers.CheckDBError(err)
		return copyCount
	} else {
		return 0
	}
}
