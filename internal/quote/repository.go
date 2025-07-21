package quote

import (
	"context"
	"frete-rapido/internal/entity"
)

type Repository interface {
	SaveQuoteRequest(ctx context.Context, zipcode string) (int, error)
	SaveQuote(ctx context.Context, quoteReqID int, resp entity.QuoteResponse) error
	GetMetrics(ctx context.Context, ids []int) ([]entity.Quote, error)
	GetLastQuotes(ctx context.Context, n int) ([]int, error)
}
