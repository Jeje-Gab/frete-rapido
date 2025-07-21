package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type QuoteRequestRepo struct {
	db *sql.DB
}

func NewQuoteRequestRepo(db *sql.DB) *QuoteRequestRepo {
	return &QuoteRequestRepo{db: db}
}

// Cria uma nova quote request e retorna o ID criado usado para filtros em metrics
func (r *QuoteRequestRepo) Execute(ctx context.Context, zipcode string) (int, error) {
	var id int
	err := r.db.QueryRowContext(
		ctx,
		queryInsertQuoteRequest,
		zipcode,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error inserting quote request: %w", err)
	}
	return id, nil
}

var (
	queryInsertQuoteRequest = `
		INSERT INTO frete.quote_requests (recipient_zipcode) 
		VALUES ($1) 
		RETURNING id
	`
)
