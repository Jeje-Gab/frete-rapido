package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type GetLastQuotesRepo struct {
	db *sql.DB
}

// NewGetLastQuotesRepository inicializa o repositório.
func NewGetLastQuotesRepository(db *sql.DB) *GetLastQuotesRepo {
	return &GetLastQuotesRepo{db: db}
}

// retorna as últimas cotações ordenadas da mais recente para a mais antiga.
func (r *GetLastQuotesRepo) Execute(ctx context.Context, n int) ([]int, error) {
	var (
		rows *sql.Rows
		err  error
	)

	switch {
	case n > 0:
		rows, err = r.db.QueryContext(ctx, queryWithLimit, n)
	default:
		rows, err = r.db.QueryContext(ctx, queryWithoutLimit)
	}
	if err != nil {
		return nil, fmt.Errorf("error querying quote_requests: %w", err)
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	var ids []int
	for rows.Next() {
		var id int
		if scanErr := rows.Scan(&id); scanErr != nil {
			return nil, fmt.Errorf("error scanning row: %w", scanErr)
		}
		ids = append(ids, id)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	if len(ids) == 0 {
		return nil, errors.New("no quote_requests found")
	}
	return ids, nil
}

var (
	queryWithLimit = `
		SELECT id 
		FROM frete.quote_requests 
		ORDER BY created_at DESC LIMIT $1
	`

	queryWithoutLimit = `
		SELECT id 
		FROM frete.quote_requests 
		ORDER BY created_at DESC
	`
)
