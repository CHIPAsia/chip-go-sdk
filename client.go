package chip

type Client struct {
	Description           string   `json:"description"`
	Email                 string   `json:"email"`
	Phone                 string   `json:"phone"`
	Fullname              string   `json:"full_name"`
	PersonalCode          string   `json:"personal_code"`
	StreetAddress         string   `json:"street_address"`
	Country               string   `json:"country"`
	City                  string   `json:"city"`
	ZipCode               string   `json:"zip_code"`
	State                 string   `json:"state"`
	ShippingStreetAddress string   `json:"shipping_street_address"`
	ShippingCountry       string   `json:"shipping_country"`
	ShippingCity          string   `json:"shipping_city"`
	ShippingZipCode       string   `json:"shipping_zip_code"`
	ShippingState         string   `json:"shipping_state"`
	CC                    []string `json:"cc"`
	BCC                   []string `json:"bcc"`
	LegalName             string   `json:"legal_name"`
	BrandName             string   `json:"brand_name"`
	RegistrationNumber    string   `json:"registration_number"`
	TaxNumber             string   `json:"tax_number"`
}
