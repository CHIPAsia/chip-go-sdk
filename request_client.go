package chip

type RequestClient string

const (
	EMAIL               RequestClient = "email"
	PHONE               RequestClient = "phone"
	FULLNAME            RequestClient = "full_name"
	PERSONAL_CODE       RequestClient = "personal_code"
	BRAND_NAME          RequestClient = "brand_name"
	LEGAL_NAME          RequestClient = "legal_name"
	REGISTRATION_NUMBER RequestClient = "registration_number"
	TAX_NUMBER          RequestClient = "tax_number"
	BANK_ACCOUNT        RequestClient = "bank_account"
	BANK_CODE           RequestClient = "bank_code"
	BILLING_ADDERESS    RequestClient = "billing_address"
	SHIPPING_ADDRESS    RequestClient = "shipping_address"
)
