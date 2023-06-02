package chip

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/CHIPAsia/chip-go-sdk/internal/http"
)

type paymentMethodsInterface interface {
	setHttpClient(http.HttpRequest)
	// GetMethods returns available methods by Brand ID
	GetMethods(PaymentMethodOption) (PaymentMethod, error)
}

type paymentMethods struct {
	accessToken string
	brandID     string
	basePath    string
	http        http.HttpRequest
}

func newPaymentMethods(accessToken, brandID, basePath string) paymentMethodsInterface {
	return &paymentMethods{
		accessToken: accessToken,
		brandID:     brandID,
		basePath:    basePath,
		http:        http.DefaultHttpRequest{},
	}
}

func (pm *paymentMethods) setHttpClient(newClient http.HttpRequest) {
	pm.http = newClient
}

func (pm paymentMethods) GetMethods(options PaymentMethodOption) (PaymentMethod, error) {
	if options.BrandID == nil || len(strings.TrimSpace(*options.BrandID)) <= 0 {
		options.BrandID = &pm.brandID
	}
	queryStrings := http.GenerateQueryString(options)

	url := fmt.Sprintf("%s/v1/payment_methods/?%s", pm.basePath, queryStrings)
	body, statusCode, err := pm.http.GetJson(url, http.Header{
		"Authorization": fmt.Sprintf("Bearer %s", pm.accessToken),
	})

	result := PaymentMethod{}
	if statusCode == http.NotFound {
		err := fmt.Errorf("result cannot be found, status code %d", statusCode)
		return result, err
	}
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return result, err
	}

	return result, nil
}
