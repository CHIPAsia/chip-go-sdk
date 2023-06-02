package chip

import (
	"testing"

	"github.com/CHIPAsia/chip-go-sdk/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetPurchase(t *testing.T) {
	assert := assert.New(t)

	apiClient := NewAPIClient("<<YOUR_ACCESS_KEY>>", "<<YOUR_BRAND_KEY", "")
	apiClient.SetHTTPClient(mocks.MockHttpRequest{})

	id := "1026967d-0705-4e24-ab35-9c125e5701ea"
	response, err := apiClient.Purchases().GetPurchase(id)
	if err != nil {
		assert.Empty(err)
	}

	assert.Equal(response.ID, id, "Purchase ID should be the same.")
}

func TestCreatePurchase(t *testing.T) {
	assert := assert.New(t)

	apiClient := NewAPIClient("<<YOUR_ACCESS_KEY>>", "<<YOUR_BRAND_KEY>>", "")
	apiClient.SetHTTPClient(mocks.MockHttpRequest{})

	purchase := PurchaseOption{
		Client: Client{
			Email:    "test@test.com",
			Fullname: "John Doe",
			Phone:    "60111222333",
		},
		Purchase: PurchaseDetail{
			Products: []Product{
				{
					Name:     "Product 1",
					Quantity: 1,
					Price:    1000,
				},
			},
		},
		SuccessRedirect: "",
		FailureRedirect: "",
		CancelRedirect:  "",
		SuccessCallback: "",
	}
	response, err := apiClient.Purchases().CreatePurchase(purchase)
	if err != nil {
		assert.Empty(err)
	}
	assert.NotEmpty(response.ID)
}
