package chip

type RefundAvailability string

const (
	ALL          RefundAvailability = "all"
	FULL_ONLY    RefundAvailability = "full_only"
	PARTIAL_ONLY RefundAvailability = "partial_only"
	PIS_ALL      RefundAvailability = "pis_all"
	PIS_PARTIAL  RefundAvailability = "pis_partial"
	NONE         RefundAvailability = "none"
)
