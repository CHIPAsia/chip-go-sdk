package chip

type TransactionData struct {
	Description   string    `json:"description"`
	PaymentMethod string    `json:"payment_method"`
	Country       string    `json:"country"`
	Attempts      []Attempt `json:"attempts"`
}
