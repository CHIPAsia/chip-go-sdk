package chip

type Purchase struct {
	Description            string             `json:"description"`
	Type                   string             `json:"type"`
	ID                     string             `json:"id"`
	CreatedOn              int                `json:"created_on"`
	UpdatedOn              int                `json:"updated_on"`
	Client                 Client             `json:"client"`
	Purchase               PurchaseDetail     `json:"purchase"`
	Payment                Payment            `json:"payment"`
	IssuerDetails          IssuerDetail       `json:"issuer_details"`
	TransactionData        TransactionData    `json:"transaction_data"`
	Status                 Status             `json:"status"`
	StatusHistory          []StatusHistory    `json:"status_history"`
	ViewedOn               int                `json:"viewed_on"`
	CompanyID              string             `json:"company_id"`
	IsTest                 bool               `json:"is_test"`
	UserID                 string             `json:"user_id"`
	BrandID                string             `json:"brand_id"`
	BillingTemplateID      string             `json:"billing_template_id"`
	ClientID               string             `json:"client_id"`
	SendReceipt            bool               `json:"send_receipt"`
	IsRecurringToken       bool               `json:"is_recurring_token"`
	RecurringToken         string             `json:"recurring_token"`
	SkipCapture            bool               `json:"skip_capture"`
	ForceRecurring         bool               `json:"force_recurring"`
	ReferenceGenerated     string             `json:"reference_generated"`
	Reference              string             `json:"reference"`
	Issued                 string             `json:"issued"`
	Due                    int                `json:"due"`
	RefundAvailability     RefundAvailability `json:"refund_availability"`
	RefundableAmount       int                `json:"refundable_amount"`
	CurrencyConversion     CurrencyConversion `json:"currency_conversion"`
	PaymentMethodWhitelist []string           `json:"payment_method_whitelist"`
	SuccessRedirect        string             `json:"success_redirect"`
	FailureRedirect        string             `json:"failure_redirect"`
	CancelRedirect         string             `json:"cancel_redirect"`
	SuccessCallback        string             `json:"success_callback"`
	CreatorAgent           string             `json:"creator_agent"`
	Platform               Platform           `json:"platform"`
	Product                ProductType        `json:"product"`
	CreatedFromIP          string             `json:"created_from_ip"`
	InvoiceURL             string             `json:"invoice_url"`
	CheckoutURL            string             `json:"checkout_url"`
	DirectPostURL          string             `json:"direct_post_url"`
	MarkedAsPaid           bool               `json:"marked_as_paid"`
	OrderID                string             `json:"order_id"`
	UpsellCampaigns        []string           `json:"upsell_campaigns"`
	ReferralCampaignID     string             `json:"referral_campaign_id"`
	ReferralCode           string             `json:"referral_code"`
	ReferralCodeGenerated  string             `json:"referral_code_generated"`
	ReferralCodeDetails    ReferralCodeDetail `json:"referral_code_details"`
	RetainLevelDetails     RetainLevelDetail  `json:"retain_level_details"`
}

type PurchaseOption struct {
	Client                 Client         `json:"client"`
	Purchase               PurchaseDetail `json:"purchase"`
	BrandID                *string        `json:"brand_id"`
	ClientID               string         `json:"client_id"`
	SendReceipt            bool           `json:"send_receipt"`
	SkipCapture            bool           `json:"skip_capture"`
	ForceRecurring         bool           `json:"force_recurring"`
	Reference              string         `json:"reference"`
	Issued                 string         `json:"issued"`
	Due                    int            `json:"due"`
	PaymentMethodWhitelist []string       `json:"payment_method_whitelist"`
	SuccessRedirect        string         `json:"success_redirect"`
	FailureRedirect        string         `json:"failure_redirect"`
	CancelRedirect         string         `json:"cancel_redirect"`
	SuccessCallback        string         `json:"success_callback"`
	CreatorAgent           string         `json:"creator_agent"`
	Platform               Platform       `json:"platform"`
	UpsellCampaigns        []string       `json:"upsell_campaigns"`
	ReferralCampaignID     string         `json:"referral_campaign_id"`
}
