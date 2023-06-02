package chip

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"

	"github.com/CHIPAsia/chip-go-sdk/internal/http"
)

type purchasesInterface interface {
	setHttpClient(http.HttpRequest)
	// GetPurchase requires Purchase ID and
	// returns Purchase
	GetPurchase(id string) (Purchase, error)
	// CreatePurchase creates purchase for its Brand ID
	CreatePurchase(PurchaseOption) (Purchase, error)
	VerifyCallback(rawPayload, xSignature, publicKey string) (bool, error)
}

type purchases struct {
	accessToken string
	brandID     string
	basePath    string
	http        http.HttpRequest
}

func newPurchases(accessToken, brandID, basePath string) purchasesInterface {
	return &purchases{
		accessToken: accessToken,
		brandID:     brandID,
		basePath:    basePath,
		http:        http.DefaultHttpRequest{},
	}
}

func (p *purchases) setHttpClient(newClient http.HttpRequest) {
	p.http = newClient
}

func (p purchases) GetPurchase(id string) (Purchase, error) {
	url := fmt.Sprintf("%s/v1/purchases/%s/", p.basePath, id)
	body, statusCode, err := p.http.GetJson(url, http.Header{
		"Authorization": fmt.Sprintf("Bearer %s", p.accessToken),
	})

	result := Purchase{}
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

func (p purchases) CreatePurchase(options PurchaseOption) (Purchase, error) {
	result := Purchase{}

	if options.BrandID == nil || len(strings.TrimSpace(*options.BrandID)) <= 0 {
		options.BrandID = &p.brandID
	}
	data, err := json.Marshal(options)
	if err != nil {
		return result, err
	}
	dataString := string(data)

	url := fmt.Sprintf("%s/v1/purchases/", p.basePath)
	body, _, err := p.http.PostJson(url, http.Header{
		"Authorization": fmt.Sprintf("Bearer %s", p.accessToken),
	}, &dataString)

	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return result, err
	}

	return result, nil
}

func (p purchases) VerifyCallback(rawPayload, xSignature, publicKey string) (bool, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		err := errors.New("unable to decode public key")
		return false, err
	}

	signedBytes, err := base64.StdEncoding.DecodeString(xSignature)
	if err != nil {
		return false, err
	}

	sha256 := crypto.SHA256.New()
	sha256.Write([]byte(rawPayload))
	hashed := sha256.Sum(nil)

	pubK, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}

	if pubKey, ok := pubK.(*rsa.PublicKey); ok {
		if err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed, signedBytes); err != nil {
			return false, err
		}
		return true, nil
	}

	err = errors.New("unable to verify the signature")
	return false, err
}
