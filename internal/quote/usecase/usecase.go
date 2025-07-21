// internal/quote/usecase/usecase.go
package usecase

import (
	"context"
	"frete-rapido/internal/entity"
	"frete-rapido/internal/quote"
)

// UseCase implementa quote.UseCase (interface do domínio).
type UseCase struct {
	saveQuote *SaveQuoteUC
	getMetric *MetricsUC
}

type FreteRapidoClient interface {
	Cotar(ctx context.Context, req entity.QuoteRequest) (entity.QuoteResponse, error)
}

// NewUseCase cria o usecase com as dependências.
// Agora recebe repo E client.
func NewUseCase(quoteRepo quote.Repository, client FreteRapidoClient) quote.UseCase {
	return &UseCase{
		saveQuote: NewQuoteUseCase(quoteRepo, client),
		getMetric: NewMetricsUseCase(quoteRepo),
	}
}

func (u *UseCase) Cotar(ctx context.Context, req entity.QuoteRequest) (entity.QuoteResponse, error) {
	return u.saveQuote.Execute(ctx, req)
}

func (u *UseCase) GetMetrics(ctx context.Context, lastQuotes int) (entity.MetricsSummary, error) {
	return u.getMetric.Execute(ctx, lastQuotes)
}
