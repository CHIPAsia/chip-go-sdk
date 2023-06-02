package chip

type Payment struct {
	Description       string `json:"description"`
	IsOutgoing        bool   `json:"is_outgoing"`
	PaymentType       string `json:"payment_type"`
	Amount            int    `json:"amount"`
	Currency          string `json:"currency"`
	NetAmount         int    `json:"net_amount"`
	FeeAmount         int    `json:"fee_amount"`
	PendingAmount     int    `json:"pending_amount"`
	PendingUnfreezeOn int    `json:"pending_unfreeze_on"`
	PaidOn            int    `json:"paid_on"`
	RemotePaidOn      int    `json:"remove_paid_on"`
}
