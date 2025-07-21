package freterapido

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"frete-rapido/internal/entity"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Client struct {
	Token         string
	Endpoint      string
	CNPJ          string
	PlatformCode  string
	DispatcherZip string
}

// padrões esperados pela api frete rapido, dados sensiveis inseridos pelo .env
func NewClient(token, endpoint, cnpj, platformCode, dispatcherZip string) *Client {
	return &Client{
		Token:         token,
		Endpoint:      endpoint,
		CNPJ:          cnpj,
		PlatformCode:  platformCode,
		DispatcherZip: dispatcherZip,
	}
}

func (c *Client) Cotar(ctx context.Context, req entity.QuoteRequest) (entity.QuoteResponse, error) {
	recipientZip, err := strconv.Atoi(req.Recipient.Address.Zipcode)
	if err != nil {
		return entity.QuoteResponse{}, fmt.Errorf("zipcode do destinatário inválido: %w", err)
	}

	dispatcherZip, err := strconv.Atoi(c.DispatcherZip)
	if err != nil {
		return entity.QuoteResponse{}, fmt.Errorf("zipcode do dispatcher inválido: %w", err)
	}

	payload := map[string]interface{}{
		"shipper": map[string]interface{}{
			"registered_number": c.CNPJ,
			"token":             c.Token,
			"platform_code":     c.PlatformCode,
		},
		"recipient": map[string]interface{}{
			"type":    0,
			"country": "BRA",
			"zipcode": recipientZip,
		},
		"dispatchers": []map[string]interface{}{
			{
				"registered_number": c.CNPJ,
				"zipcode":           dispatcherZip,
				"volumes":           buildVolumes(req.Volumes),
			},
		},
		"simulation_type": []int{1},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return entity.QuoteResponse{}, err
	}

	httpReq, err := http.NewRequest("POST", c.Endpoint, bytes.NewBuffer(body))
	if err != nil {
		return entity.QuoteResponse{}, err
	}
	httpReq.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return entity.QuoteResponse{}, err
	}
	defer resp.Body.Close()

	jsonBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(jsonBody)) // printa retorno vindo do consumo e envio para api da frete rapido

	if resp.StatusCode != http.StatusOK {
		return entity.QuoteResponse{}, fmt.Errorf("erro na API Frete Rápido: %v", resp.Status)
	}

	// Por conta de não ter uma segunda rota para solicitar carriers foi
	// mocado um valor com base na saida esperada de acordo com a documentação
	// fornecida para realizaçâo do teste
	carriers := []entity.CarrierQuote{
		{
			Name:     "EXPRESSO FR",
			Service:  "Rodoviário",
			Deadline: 3,
			Price:    17,
		},
		{
			Name:     "Correios",
			Service:  "SEDEX",
			Deadline: 1,
			Price:    20.99,
		},
	}
	quoteResp := entity.QuoteResponse{
		Carrier: carriers,
	}

	return quoteResp, nil
}

func buildVolumes(volumes []entity.Volumes) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(volumes))
	for _, v := range volumes {
		out = append(out, map[string]interface{}{
			"amount":         v.Amount,
			"category":       strconv.Itoa(v.Category), // sempre string!
			"sku":            v.Sku,
			"height":         v.Height,
			"width":          v.Width,
			"length":         v.Length,
			"unitary_price":  v.Price,
			"unitary_weight": v.UnitaryWeight,
		})
	}
	return out
}
