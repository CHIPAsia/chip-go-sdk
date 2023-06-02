package chip

type ProductType string

const (
	PURCHASES                     ProductType = "purchases"
	BILLING_INVOICES              ProductType = "billing_invoices"
	BILLING_SUBSCRIPTIONS         ProductType = "billing_subscriptions"
	BILLING_SUBSCRIPTIONS_INVOICE ProductType = "billing_subscriptions_invoice"
)
