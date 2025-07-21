package repository

import (
	"context"
	"database/sql"
	"frete-rapido/internal/entity"
	"frete-rapido/internal/quote"
)

// quoteRepo implementa negociacoes.Repository usando PostgreSQL.
type quoteRepo struct {
	saveQuote        *SaveQuoteRepo
	getLastQuotes    *GetLastQuotesRepo
	getMetrics       *GetMetricsRepo
	saveQuoteRequest *QuoteRequestRepo
}

func NewRepository(db *sql.DB) quote.Repository {
	return &quoteRepo{
		saveQuote:        NewSaveQuoteRepository(db),
		getLastQuotes:    NewGetLastQuotesRepository(db),
		getMetrics:       NewGetMetricsRepository(db),
		saveQuoteRequest: NewQuoteRequestRepo(db),
	}
}

func (r *quoteRepo) SaveQuote(ctx context.Context, quoteRequestID int, resp entity.QuoteResponse) error {
	return r.saveQuote.Execute(ctx, quoteRequestID, resp)
}

func (r *quoteRepo) GetMetrics(ctx context.Context, ids []int) ([]entity.Quote, error) {
	return r.getMetrics.Execute(ctx, ids)
}

func (r *quoteRepo) SaveQuoteRequest(ctx context.Context, zipcode string) (int, error) {
	return r.saveQuoteRequest.Execute(ctx, zipcode)
}

func (r *quoteRepo) GetLastQuotes(ctx context.Context, n int) ([]int, error) {
	return r.getLastQuotes.Execute(ctx, n)
}
