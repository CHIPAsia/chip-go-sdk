package chip

import (
	"testing"

	"github.com/CHIPAsia/chip-go-sdk/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetMethods(t *testing.T) {
	assert := assert.New(t)

	apiClient := NewAPIClient("<<YOUR_ACCESS_KEY>>", "<<YOUR_BRAND_KEY", "")
	apiClient.SetHTTPClient(mocks.MockHttpRequest{})

	options := PaymentMethodOption{
		Currency: "MYR",
	}
	response, err := apiClient.PaymentMethods().GetMethods(options)
	if err != nil {
		assert.Empty(err)
	}

	assert.NotEmpty(response.AvailablePaymentMethods)
}
