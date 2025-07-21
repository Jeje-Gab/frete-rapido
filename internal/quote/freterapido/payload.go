package freterapido

type FRQuoteRequest struct {
	Shipper struct {
		RegisteredNumber string `json:"registered_number"`
	} `json:"shipper"`
	Dispatchers []struct {
		RegisteredNumber string `json:"registered_number"`
		Zipcode          string `json:"zipcode"`
	} `json:"dispatchers"`
	Recipient struct {
		Address struct {
			Zipcode string `json:"zipcode"`
		} `json:"address"`
	} `json:"recipient"`
	Volumes      []Volume `json:"volumes"`
	PlatformCode string   `json:"platform_code"`
}

type Volume struct {
	Category      int     `json:"category"`
	Amount        int     `json:"amount"`
	UnitaryWeight float64 `json:"unitary_weight"`
	Price         float64 `json:"price"`
	SKU           string  `json:"sku"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
}
