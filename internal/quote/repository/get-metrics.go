package repository

import (
	"context"
	"database/sql"
	"fmt"

	"frete-rapido/internal/entity"
	"github.com/lib/pq"
)

type GetMetricsRepo struct {
	db *sql.DB
}

func NewGetMetricsRepository(db *sql.DB) *GetMetricsRepo {
	return &GetMetricsRepo{db: db}
}

// retorna uma lista de Quotes para os IDs informados.
func (r *GetMetricsRepo) Execute(ctx context.Context, ids []int) ([]entity.Quote, error) {
	if len(ids) == 0 {
		return []entity.Quote{}, nil
	}

	rows, err := r.db.QueryContext(ctx, querySelectQuotesByIDs, pq.Array(ids))
	if err != nil {
		return nil, fmt.Errorf("error querying quotes: %w", err)
	}
	defer rows.Close()

	var quotes []entity.Quote
	for rows.Next() {
		var q entity.Quote
		if err := rows.Scan(&q.CarrierName, &q.Service, &q.Deadline, &q.Price); err != nil {
			return nil, fmt.Errorf("error scanning quote row: %w", err)
		}
		quotes = append(quotes, q)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return quotes, nil
}

var (
	querySelectQuotesByIDs = `
		SELECT carrier_name, service, deadline, price 
		FROM frete.quotes 
		WHERE quote_request_id = ANY($1)
	`
)
