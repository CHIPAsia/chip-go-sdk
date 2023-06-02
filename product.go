package chip

type Product struct {
	Name       string      `json:"name"`
	Quantity   interface{} `json:"quantity"`
	Price      int         `json:"price"`
	Discount   int         `json:"discount"`
	TaxPercent string      `json:"tax_percent"`
	Category   string      `json:"category"`
}
