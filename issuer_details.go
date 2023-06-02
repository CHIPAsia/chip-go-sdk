package chip

type IssuerDetail struct {
	Description        string        `json:"description"`
	Website            string        `json:"website"`
	LegalStreetAddress string        `json:"legal_street_address"`
	LegalCountry       string        `json:"legal_country"`
	LegalCity          string        `json:"legal_city"`
	LegalZipCode       string        `json:"legal_zip_code"`
	BankAccounts       []BankAccount `json:"bank_accounts"`
	LegalName          string        `json:"legal_name"`
	BrandName          string        `json:"brand_name"`
	RegistrationNumber string        `json:"registration_number"`
	TaxNumber          string        `json:"tax_number"`
}
