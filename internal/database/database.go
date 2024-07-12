package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"example.com/mamude/internal/helpers"
	"example.com/mamude/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	// Save Customers
	SaveCustomers(c *gin.Context, customers []types.Customer) int64
}

type service struct {
	db *pgx.Conn
}

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	schema     = os.Getenv("DB_SCHEMA")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
	db, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}
	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.Ping(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf(fmt.Sprintf("db down: %v", err)) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"
	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	return s.db.Close(ctx)
}

// Salvar dados tratados do cliente no DB
func (s *service) SaveCustomers(ctx *gin.Context, customers []types.Customer) int64 {
	// tratar os dados e validá-los
	var customersValid []types.Customer
	for _, customer := range customers {
		customer.ValidateCPF(customer.CPF)
		customer.ValidateMostFrequentStore(customer.MostFrequentStore)
		customer.ValidateLastPurchaseStore(customer.LastPurchaseStore)
		customersValid = append(customersValid, customer)
	}

	// enviar os dados em batch, otimizando o desempenho
	// https://pkg.go.dev/github.com/jackc/pgx/v5@v5.3.0#hdr-Copy_Protocol
	copyCount, err := s.db.CopyFrom(
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

	// tx, err := s.db.Begin()
	// helpers.CheckDBError(err)
	// defer tx.Rollback()

	// sql, err := s.db.Prepare(`
	// 	INSERT INTO customers
	// 	(cpf, private, incomplete, date_last_purchase, average_ticket, last_purchase_ticket, most_frequent_store, last_purchase_store)
	// 	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	// `)
	// helpers.CheckDBError(err)
	// defer sql.Close()

	// for idx, customer := range customers {
	// 	log.Printf("Line: %v", idx)

	// 	// tratar e validar documentos
	// 	customer.ValidateCPF(customer.CPF)
	// 	customer.ValidateMostFrequentStore(customer.MostFrequentStore)
	// 	customer.ValidateLastPurchaseStore(customer.LastPurchaseStore)

	// 	if _, err := sql.Exec(
	// 		customer.CPF,
	// 		customer.Private,
	// 		customer.Incomplete,
	// 		customer.DateLastPurchase,
	// 		customer.AverageTicket,
	// 		customer.LastPurchaseTicket,
	// 		customer.MostFrequentStore,
	// 		customer.LastPurchaseStore,
	// 	); err != nil {
	// 		panic(err)
	// 	}
	// }

	// if err := tx.Commit(); err != nil {
	// 	panic(err)
	// }
}
