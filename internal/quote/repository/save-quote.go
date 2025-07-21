package repository

import (
	"context"
	"database/sql"
	"fmt"

	"frete-rapido/internal/entity"
)

type SaveQuoteRepo struct {
	db *sql.DB
}

func NewSaveQuoteRepository(db *sql.DB) *SaveQuoteRepo {
	return &SaveQuoteRepo{db: db}
}

func (r *SaveQuoteRepo) Execute(ctx context.Context, quoteRequestID int, resp entity.QuoteResponse) error {
	for _, carrier := range resp.Carrier {
		_, err := r.db.ExecContext(
			ctx,
			queryInsertQuote,
			quoteRequestID,
			carrier.Name,
			carrier.Service,
			carrier.Deadline,
			carrier.Price,
		)
		if err != nil {
			return fmt.Errorf("error inserting quote: %w", err)
		}
	}
	return nil
}

var (
	queryInsertQuote = `
		INSERT INTO frete.quotes
			(quote_request_id, carrier_name, service, deadline, price)
		VALUES ($1, $2, $3, $4, $5)
	`
)
