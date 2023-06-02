package chip

type Status string

const (
	CREATED         Status = "created"
	SENT            Status = "sent"
	VIEWED          Status = "viewed"
	ERROR           Status = "error"
	CANCELLED       Status = "cancelled"
	OVERDUE         Status = "overdue"
	EXPIRED         Status = "expired"
	BLOCKED         Status = "blocked"
	HOLD            Status = "hold"
	RELEASED        Status = "released"
	PENDING_RELEASE Status = "pending_release"
	PENDING_CAPTURE Status = "pending_capture"
	PREAUTHORIZED   Status = "preauthorized"
	PAID            Status = "paid"
	PENDING_EXECUTE Status = "pending_execute"
	PENDING_CHARGE  Status = "pending_charge"
	CLEARED         Status = "cleared"
	SETTLED         Status = "settled"
	CHARGEBACK      Status = "chargeback"
	PENDING_REFUND  Status = "pending_refund"
	REFUNDED        Status = "refunded"
)
