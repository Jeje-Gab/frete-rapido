// internal/quote/usecase/save-quote.go
package usecase

import (
	"context"
	"frete-rapido/internal/entity"
	"frete-rapido/internal/quote"
)

type SaveQuoteUC struct {
	quoteRepo quote.Repository
	client    FreteRapidoClient
}

func NewQuoteUseCase(quoteRepo quote.Repository, client FreteRapidoClient) *SaveQuoteUC {
	return &SaveQuoteUC{quoteRepo: quoteRepo, client: client}
}

func (u *SaveQuoteUC) Execute(ctx context.Context, req entity.QuoteRequest) (entity.QuoteResponse, error) {
	quoteReqID, err := u.quoteRepo.SaveQuoteRequest(ctx, req.Recipient.Address.Zipcode)
	if err != nil {
		return entity.QuoteResponse{}, err
	}

	// Chama o client da Frete Rápido inicia a conexão e faz mock do retorno
	resp, err := u.client.Cotar(ctx, req)
	if err != nil {
		return entity.QuoteResponse{}, err
	}

	if err := u.quoteRepo.SaveQuote(ctx, quoteReqID, resp); err != nil {
		return entity.QuoteResponse{}, err
	}

	return resp, nil
}
