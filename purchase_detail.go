package chip

type PurchaseDetail struct {
	Description           string        `json:"description"`
	Currency              string        `json:"currency"`
	Products              []Product     `json:"products"`
	Total                 int           `json:"total"`
	Language              string        `json:"language"`
	Notes                 string        `json:"notes"`
	Debt                  int           `json:"debt"`
	SubtotalOverride      int           `json:"subtotal_override"`
	TotalTaxOverride      int           `json:"total_tax_override"`
	TotalDiscountOverride int           `json:"total_discount_override"`
	TotalOverride         int           `json:"total_override"`
	RequestClientDetail   RequestClient `json:"request_client_detail"`
	Timezone              string        `json:"timezone"`
	DueStrict             bool          `json:"due_strict"`
	EmalMessage           string        `json:"email_string"`
}
