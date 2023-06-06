package chip

type Purchase struct {
	Description            string             `json:"description,omitempty"`
	Type                   string             `json:"type,omitempty"`
	ID                     string             `json:"id,omitempty"`
	CreatedOn              int                `json:"created_on,omitempty"`
	UpdatedOn              int                `json:"updated_on,omitempty"`
	Client                 Client             `json:"client,omitempty"`
	Purchase               PurchaseDetail     `json:"purchase,omitempty"`
	Payment                Payment            `json:"payment,omitempty"`
	IssuerDetails          IssuerDetail       `json:"issuer_details,omitempty"`
	TransactionData        TransactionData    `json:"transaction_data,omitempty"`
	Status                 Status             `json:"status,omitempty"`
	StatusHistory          []StatusHistory    `json:"status_history,omitempty"`
	ViewedOn               int                `json:"viewed_on,omitempty"`
	CompanyID              string             `json:"company_id,omitempty"`
	IsTest                 bool               `json:"is_test,omitempty"`
	UserID                 string             `json:"user_id,omitempty"`
	BrandID                string             `json:"brand_id,omitempty"`
	BillingTemplateID      string             `json:"billing_template_id,omitempty"`
	ClientID               string             `json:"client_id,omitempty"`
	SendReceipt            bool               `json:"send_receipt,omitempty"`
	IsRecurringToken       bool               `json:"is_recurring_token,omitempty"`
	RecurringToken         string             `json:"recurring_token,omitempty"`
	SkipCapture            bool               `json:"skip_capture,omitempty"`
	ForceRecurring         bool               `json:"force_recurring,omitempty"`
	ReferenceGenerated     string             `json:"reference_generated,omitempty"`
	Reference              string             `json:"reference,omitempty"`
	Issued                 string             `json:"issued,omitempty"`
	Due                    int                `json:"due,omitempty"`
	RefundAvailability     RefundAvailability `json:"refund_availability,omitempty"`
	RefundableAmount       int                `json:"refundable_amount,omitempty"`
	CurrencyConversion     CurrencyConversion `json:"currency_conversion,omitempty"`
	PaymentMethodWhitelist []string           `json:"payment_method_whitelist,omitempty"`
	SuccessRedirect        string             `json:"success_redirect,omitempty"`
	FailureRedirect        string             `json:"failure_redirect,omitempty"`
	CancelRedirect         string             `json:"cancel_redirect,omitempty"`
	SuccessCallback        string             `json:"success_callback,omitempty"`
	CreatorAgent           string             `json:"creator_agent,omitempty"`
	Platform               Platform           `json:"platform,omitempty"`
	Product                ProductType        `json:"product,omitempty"`
	CreatedFromIP          string             `json:"created_from_ip,omitempty"`
	InvoiceURL             string             `json:"invoice_url,omitempty"`
	CheckoutURL            string             `json:"checkout_url,omitempty"`
	DirectPostURL          string             `json:"direct_post_url,omitempty"`
	MarkedAsPaid           bool               `json:"marked_as_paid,omitempty"`
	OrderID                string             `json:"order_id,omitempty"`
	UpsellCampaigns        []string           `json:"upsell_campaigns,omitempty"`
	ReferralCampaignID     string             `json:"referral_campaign_id,omitempty"`
	ReferralCode           string             `json:"referral_code,omitempty"`
	ReferralCodeGenerated  string             `json:"referral_code_generated,omitempty"`
	ReferralCodeDetails    ReferralCodeDetail `json:"referral_code_details,omitempty"`
	RetainLevelDetails     RetainLevelDetail  `json:"retain_level_details,omitempty"`
}

type PurchaseOption struct {
	Client                 Client         `json:"client,omitempty"`
	Purchase               PurchaseDetail `json:"purchase,omitempty"`
	BrandID                *string        `json:"brand_id,omitempty"`
	ClientID               string         `json:"client_id,omitempty"`
	SendReceipt            bool           `json:"send_receipt,omitempty"`
	SkipCapture            bool           `json:"skip_capture,omitempty"`
	ForceRecurring         bool           `json:"force_recurring,omitempty"`
	Reference              string         `json:"reference,omitempty"`
	Issued                 string         `json:"issued,omitempty"`
	Due                    int            `json:"due,omitempty"`
	PaymentMethodWhitelist []string       `json:"payment_method_whitelist,omitempty"`
	SuccessRedirect        string         `json:"success_redirect,omitempty"`
	FailureRedirect        string         `json:"failure_redirect,omitempty"`
	CancelRedirect         string         `json:"cancel_redirect,omitempty"`
	SuccessCallback        string         `json:"success_callback,omitempty"`
	CreatorAgent           string         `json:"creator_agent,omitempty"`
	Platform               Platform       `json:"platform,omitempty"`
	UpsellCampaigns        []string       `json:"upsell_campaigns,omitempty"`
	ReferralCampaignID     string         `json:"referral_campaign_id,omitempty"`
}
