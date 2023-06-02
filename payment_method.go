package chip

type PaymentMethod struct {
	AvailablePaymentMethods []string               `json:"available_payment_methods"`
	ByCountry               map[string][]string    `json:"by_country"`
	CountryNames            map[string]string      `json:"country_names"`
	Names                   map[string]string      `json:"names"`
	Logos                   map[string]interface{} `json:"logos"`
	CardMethods             []string               `json:"card_methods"`
}

type PaymentMethodOption struct {
	BrandID          *string `json:"brand_id"`
	Currency         string  `json:"currency"`
	Country          string  `json:"country"`
	Recurring        *bool   `json:"recurring"`
	SkipCapture      *bool   `json:"skip_capture"`
	PreAuthorization *bool   `json:"preauthorization"`
	Language         string  `json:"language"`
}
