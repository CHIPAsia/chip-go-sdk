package chip

type PurchaseDetail struct {
	Description           string        `json:"description,omitempty"`
	Currency              string        `json:"currency,omitempty"`
	Products              []Product     `json:"products,omitempty"`
	Total                 int           `json:"total,omitempty"`
	Language              string        `json:"language,omitempty"`
	Notes                 string        `json:"notes,omitempty"`
	Debt                  int           `json:"debt,omitempty"`
	SubtotalOverride      int           `json:"subtotal_override,omitempty"`
	TotalTaxOverride      int           `json:"total_tax_override,omitempty"`
	TotalDiscountOverride int           `json:"total_discount_override,omitempty"`
	TotalOverride         int           `json:"total_override,omitempty"`
	RequestClientDetail   RequestClient `json:"request_client_detail,omitempty"`
	Timezone              string        `json:"timezone,omitempty"`
	DueStrict             bool          `json:"due_strict,omitempty"`
	EmalMessage           string        `json:"email_string,omitempty"`
}
