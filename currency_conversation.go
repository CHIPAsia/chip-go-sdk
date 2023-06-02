package chip

type CurrencyConversion struct {
	Description      string  `json:"description"`
	OriginalCurrency string  `json:"original_currency"`
	OriginalAmount   int     `json:"original_amount"`
	ExchangeRate     float32 `json:"exchange_rate"`
}
