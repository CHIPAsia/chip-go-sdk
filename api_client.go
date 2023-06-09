package chip

import (
	"strings"

	"github.com/CHIPAsia/chip-go-sdk/internal/http"
)

type APIClient interface {
	PaymentMethods() paymentMethodsInterface
	Purchases() purchasesInterface
	PublicKeys() purchasesInterface
	SetHTTPClient(http.HttpRequest)
}

type apiClient struct {
	accessToken    string
	brandID        string
	basePath       string
	paymentMethods paymentMethodsInterface
	purchases      purchasesInterface
	publicKeys     publicKeyInterface
}

func NewAPIClient(accessToken, brandID, basePath string) APIClient {
	_basePath := "https://gate.chip-in.asia/api"
	if len(strings.TrimSpace(basePath)) > 0 {
		_basePath = basePath
	}

	paymentMethods := newPaymentMethods(accessToken, brandID, _basePath)
	purchases := newPurchases(accessToken, brandID, _basePath)
	publicKeys := newPublicKey(accessToken, brandID, _basePath)
	return apiClient{
		accessToken:    accessToken,
		brandID:        brandID,
		basePath:       _basePath,
		paymentMethods: paymentMethods,
		purchases:      purchases,
		publicKeys:     publicKeys,
	}
}

func (c apiClient) PaymentMethods() paymentMethodsInterface {
	return c.paymentMethods
}

func (c apiClient) Purchases() purchasesInterface {
	return c.purchases
}

func (c apiClient) PublicKeys() purchasesInterface {
	return c.purchases
}

func (c apiClient) SetHTTPClient(newClient http.HttpRequest) {
	c.paymentMethods.setHttpClient(newClient)
	c.purchases.setHttpClient(newClient)
	c.publicKeys.setHttpClient(newClient)
}
