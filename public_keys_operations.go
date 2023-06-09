package chip

import (
	"encoding/json"
	"fmt"

	"github.com/CHIPAsia/chip-go-sdk/internal/http"
)

type publicKeysInterface interface {
	setHttpClient(http.HttpRequest)
	// GetPublicKey returns public key for the account
	GetPublicKey() (string, error)
}

type publicKeys struct {
	accessToken string
	brandID     string
	basePath    string
	http        http.HttpRequest
}

func newPublicKey(accessToken, brandID, basePath string) publicKeysInterface {
	return &publicKeys{
		accessToken: accessToken,
		brandID:     brandID,
		basePath:    basePath,
		http:        http.DefaultHttpRequest{},
	}
}

func (p *publicKeys) setHttpClient(newClient http.HttpRequest) {
	p.http = newClient
}

func (p publicKeys) GetPublicKey() (string, error) {
	url := fmt.Sprintf("%s/v1/public_key/", p.basePath)
	body, statusCode, err := p.http.GetJson(url, http.Header{
		"Authorization": fmt.Sprintf("Bearer %s", p.accessToken),
	})

	result := ""
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
