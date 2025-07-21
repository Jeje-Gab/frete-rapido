// internal/quote/usecase.go
package quote

import (
	"context"
	"frete-rapido/internal/entity"
)

// Interface para consumir a Frete Rápido
type UseCase interface {
	Cotar(ctx context.Context, req entity.QuoteRequest) (entity.QuoteResponse, error)
	GetMetrics(ctx context.Context, lastQuotes int) (entity.MetricsSummary, error)
}
