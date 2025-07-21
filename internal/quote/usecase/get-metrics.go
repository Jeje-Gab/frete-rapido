// quote/usecase/metrics.go
package usecase

import (
	"context"
	"frete-rapido/internal/entity"
	"frete-rapido/internal/quote"
)

type MetricsUC struct {
	quoteRepo quote.Repository
}

func NewMetricsUseCase(repo quote.Repository) *MetricsUC {
	return &MetricsUC{quoteRepo: repo}
}

func (u *MetricsUC) Execute(ctx context.Context, lastQuotes int) (entity.MetricsSummary, error) {

	ids, err := u.quoteRepo.GetLastQuotes(ctx, lastQuotes)
	if err != nil {
		return entity.MetricsSummary{}, err
	}

	resp, err := u.quoteRepo.GetMetrics(ctx, ids)
	if err != nil {
		return entity.MetricsSummary{}, err
	}

	return entity.CalculateMetricsSummary(resp), nil
}
