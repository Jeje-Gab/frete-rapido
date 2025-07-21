package entity

type Volumes struct {
	Category      int     `json:"category"`
	Amount        int     `json:"amount"`
	UnitaryWeight float64 `json:"unitary_weight"`
	Price         float64 `json:"price"`
	Sku           string  `json:"sku"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
}

type VolumesArray []Volumes

type Address struct {
	Zipcode string `json:"zipcode"`
}

type Recipient struct {
	Address Address `json:"address"`
}

type QuoteRequest struct {
	Recipient Recipient    `json:"recipient"`
	Volumes   VolumesArray `json:"volumes"`
}

type CarrierQuote struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline int     `json:"deadline"`
	Price    float64 `json:"price"`
}

type QuoteResponse struct {
	Carrier []CarrierQuote `json:"carrier"`
}

// --

type Quote struct {
	CarrierName string
	Service     string
	Deadline    int
	Price       float64
}

type MetricsResult struct {
	CarrierName  string  `json:"carrier_name"`
	TotalQuotes  int     `json:"total_quotes"`
	TotalPrice   float64 `json:"total_price"`
	AveragePrice float64 `json:"average_price"`
}

type MetricsSummary struct {
	ByCarrier        []MetricsResult `json:"by_carrier"`
	CheapestFreight  float64         `json:"cheapest_freight"`
	ExpensiveFreight float64         `json:"expensive_freight"`
}

func CalculateMetricsSummary(quotes []Quote) MetricsSummary {
	carrierMap := make(map[string][]float64)
	var cheapest, expensive float64
	for i, q := range quotes {
		carrierMap[q.CarrierName] = append(carrierMap[q.CarrierName], q.Price)
		if i == 0 || q.Price < cheapest {
			cheapest = q.Price
		}
		if i == 0 || q.Price > expensive {
			expensive = q.Price
		}
	}

	var byCarrier []MetricsResult
	for carrier, prices := range carrierMap {
		sum := 0.0
		for _, p := range prices {
			sum += p
		}
		byCarrier = append(byCarrier, MetricsResult{
			CarrierName:  carrier,
			TotalQuotes:  len(prices),
			TotalPrice:   sum,
			AveragePrice: sum / float64(len(prices)),
		})
	}

	return MetricsSummary{
		ByCarrier:        byCarrier,
		CheapestFreight:  cheapest,
		ExpensiveFreight: expensive,
	}
}
