package chip

type AttemptType string

const (
	EXECUTE                AttemptType = "execute"
	AUTHORIZE              AttemptType = "authorize"
	RELEASE                AttemptType = "release"
	CAPTURE                AttemptType = "capture"
	RECURRING_EXECUTE      AttemptType = "recurring_execute"
	DELETE_RECURRING_TOKEN AttemptType = "delete_recurring_token"
	REFUND                 AttemptType = "refund"
)

type Attempt struct {
	Type           AttemptType `json:"type"`
	Successful     bool        `json:"successful"`
	PaymentMethod  string      `json:"payment_method"`
	Extra          Extra       `json:"extra"`
	Country        string      `json:"country"`
	ClientIP       string      `json:"client_ip"`
	ProcessingTime int         `json:"processing_time"`
	Error          Error       `json:"error"`
}

type Extra struct {
	Description string `json:"description"`
}
