package chip

type Product struct {
	Name       string      `json:"name,omitempty"`
	Quantity   interface{} `json:"quantity,omitempty"`
	Price      int         `json:"price,omitempty"`
	Discount   int         `json:"discount,omitempty"`
	TaxPercent string      `json:"tax_percent,omitempty"`
	Category   string      `json:"category,omitempty"`
}
