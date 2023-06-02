package mocks

import (
	"regexp"

	"github.com/CHIPAsia/chip-go-sdk/internal/http"
)

type MockHttpRequest struct{}

func (h MockHttpRequest) PostJson(host string, headers http.Header, data *string) ([]byte, int, error) {
	response := ""
	match, _ := regexp.MatchString("/v1/purchases/", host)
	if match {
		response = `{"client":{"client_type":null,"email":"test@test.com","phone":"6011234566","full_name":"John Doe","personal_code":"","legal_name":"","brand_name":"","registration_number":"","tax_number":"","bank_account":"","bank_code":"","street_address":"","city":"","zip_code":"","country":"","state":"","shipping_street_address":"","shipping_city":"","shipping_zip_code":"","shipping_country":"","shipping_state":"","cc":[],"bcc":[],"delivery_methods":[{"method":"email","options":{}},{"method":"text_message","options":{"custom_message":""}}]},"purchase":{"currency":"MYR","products":[{"name":"product_1","price":35000,"quantity":"1.0000","discount":0,"tax_percent":"0.00","category":""}],"language":"en","notes":"product_1","debt":0,"subtotal_override":null,"total_tax_override":null,"total_discount_override":null,"total_override":null,"total":35000,"request_client_details":[],"timezone":"UTC","due_strict":false,"email_message":"","shipping_options":[],"payment_method_details":{},"has_upsell_products":false},"payment":{"is_outgoing":false,"payment_type":"purchase","amount":35000,"currency":"MYR","net_amount":31500,"fee_amount":0,"pending_amount":3500,"pending_unfreeze_on":1212312344,"description":"","paid_on":1212312344,"remote_paid_on":1212312344,"owned_bank_account_id":null,"owned_bank_account":null,"owned_bank_code":null},"issuer_details":{"brand_name":"BRAND_NAME","website":"","legal_name":"CHIP IN SDN. BHD.","registration_number":"REG_NO","tax_number":"","legal_street_address":"STREET","legal_country":"MY","legal_city":"Kuala Lumpur","legal_zip_code":"00000","bank_accounts":[{"bank_account":"12345567890","bank_code":"MBB"}]},"transaction_data":{"payment_method":"fpx","flow":"payform","extra":{},"country":"","attempts":[{"type":"execute","successful":true,"payment_method":"fpx","flow":"payform","extra":{},"country":"","client_ip":"","processing_time":1212312344,"fee_amount":3500,"error":null}]},"status":"paid","status_history":[{"status":"created","timestamp":1212312344},{"status":"viewed","timestamp":1212312344},{"status":"paid","timestamp":1212312344}],"viewed_on":1685592522,"force_recurring":false,"company_id":"<<COMPANY_ID>>","is_test":true,"user_id":null,"brand_id":"<<BRAND_ID>>","billing_template_id":null,"order_id":null,"client_id":"d585407b-e274-47e0-a761-344f240c5e45","send_receipt":false,"is_recurring_token":false,"recurring_token":null,"skip_capture":false,"reference_generated":"CH3804","reference":"","issued":"2023-06-01","due":1685596122,"refund_availability":"all","refundable_amount":35000,"currency_conversion":null,"payment_method_whitelist":null,"success_redirect":"https: //example.com/v1.0/redirect?status=success","failure_redirect":"https: //example.com/v1.0/redirect?status=fail","cancel_redirect":"https: //example.com/v1.0/redirect?status=cancel","success_callback":"https: //example.com/v1.0/callback","marked_as_paid":false,"upsell_campaigns":[],"referral_campaign_id":null,"referral_code":null,"referral_code_details":null,"referral_code_generated":null,"retain_level_details":null,"can_retrieve":false,"can_chargeback":false,"creator_agent":"","platform":"api","product":"purchases","created_from_ip":"","invoice_url":null,"checkout_url":"","direct_post_url":null,"created_on":1212312344,"updated_on":1212312344,"type":"purchase","id":"1026967d-0705-4e24-ab35-9c125e5701ea"}`
	}
	return []byte(response), 200, nil
}

func (h MockHttpRequest) GetJson(host string, headers http.Header) ([]byte, int, error) {
	response := ""
	match, _ := regexp.MatchString("/v1/payment_methods/", host)
	if match {
		response = `{"available_payment_methods":["visa","mastercard","some_method"],"by_country":{"any":["card"],"GB":["some_method"]},"country_names":{"any":"Other","GB":"United Kingdom"},"names":{"visa":"Visa","mastercard":"Mastercard","some_method":"Some method"},"logos":{"some_method":["/static/images/icon-visa.svg","/static/images/icon-mastercard.svg","/static/images/icon-maestro.svg"],"visa":"/static/images/icon-visa.svg","mastercard":"/static/images/icon-mastercard.svg"},"card_methods":["american_express","visa"]}`
	}
	match, _ = regexp.MatchString("/v1/purchases/[A-z0-9-]+", host)
	if match {
		response = `{"client":{"client_type":null,"email":"test@test.com","phone":"6011234566","full_name":"John Doe","personal_code":"","legal_name":"","brand_name":"","registration_number":"","tax_number":"","bank_account":"","bank_code":"","street_address":"","city":"","zip_code":"","country":"","state":"","shipping_street_address":"","shipping_city":"","shipping_zip_code":"","shipping_country":"","shipping_state":"","cc":[],"bcc":[],"delivery_methods":[{"method":"email","options":{}},{"method":"text_message","options":{"custom_message":""}}]},"purchase":{"currency":"MYR","products":[{"name":"product_1","price":35000,"quantity":"1.0000","discount":0,"tax_percent":"0.00","category":""}],"language":"en","notes":"product_1","debt":0,"subtotal_override":null,"total_tax_override":null,"total_discount_override":null,"total_override":null,"total":35000,"request_client_details":[],"timezone":"UTC","due_strict":false,"email_message":"","shipping_options":[],"payment_method_details":{},"has_upsell_products":false},"payment":{"is_outgoing":false,"payment_type":"purchase","amount":35000,"currency":"MYR","net_amount":31500,"fee_amount":0,"pending_amount":3500,"pending_unfreeze_on":1212312344,"description":"","paid_on":1212312344,"remote_paid_on":1212312344,"owned_bank_account_id":null,"owned_bank_account":null,"owned_bank_code":null},"issuer_details":{"brand_name":"BRAND_NAME","website":"","legal_name":"CHIP IN SDN. BHD.","registration_number":"REG_NO","tax_number":"","legal_street_address":"STREET","legal_country":"MY","legal_city":"Kuala Lumpur","legal_zip_code":"00000","bank_accounts":[{"bank_account":"12345567890","bank_code":"MBB"}]},"transaction_data":{"payment_method":"fpx","flow":"payform","extra":{},"country":"","attempts":[{"type":"execute","successful":true,"payment_method":"fpx","flow":"payform","extra":{},"country":"","client_ip":"","processing_time":1212312344,"fee_amount":3500,"error":null}]},"status":"paid","status_history":[{"status":"created","timestamp":1212312344},{"status":"viewed","timestamp":1212312344},{"status":"paid","timestamp":1212312344}],"viewed_on":1685592522,"force_recurring":false,"company_id":"<<COMPANY_ID>>","is_test":true,"user_id":null,"brand_id":"<<BRAND_ID>>","billing_template_id":null,"order_id":null,"client_id":"d585407b-e274-47e0-a761-344f240c5e45","send_receipt":false,"is_recurring_token":false,"recurring_token":null,"skip_capture":false,"reference_generated":"CH3804","reference":"","issued":"2023-06-01","due":1685596122,"refund_availability":"all","refundable_amount":35000,"currency_conversion":null,"payment_method_whitelist":null,"success_redirect":"https: //example.com/v1.0/redirect?status=success","failure_redirect":"https: //example.com/v1.0/redirect?status=fail","cancel_redirect":"https: //example.com/v1.0/redirect?status=cancel","success_callback":"https: //example.com/v1.0/callback","marked_as_paid":false,"upsell_campaigns":[],"referral_campaign_id":null,"referral_code":null,"referral_code_details":null,"referral_code_generated":null,"retain_level_details":null,"can_retrieve":false,"can_chargeback":false,"creator_agent":"","platform":"api","product":"purchases","created_from_ip":"","invoice_url":null,"checkout_url":"","direct_post_url":null,"created_on":1212312344,"updated_on":1212312344,"type":"purchase","id":"1026967d-0705-4e24-ab35-9c125e5701ea"}`
	}

	return []byte(response), 200, nil
}

func (h MockHttpRequest) PatchJson(host string, headers http.Header, data *string) ([]byte, int, error) {
	return []byte("body"), 200, nil
}

func (h MockHttpRequest) PutJson(host string, headers http.Header, data *string) ([]byte, int, error) {
	return []byte("body"), 200, nil
}

func (h MockHttpRequest) DeleteJson(host string, headers http.Header) ([]byte, int, error) {
	return []byte("body"), 200, nil
}
