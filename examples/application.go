package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	chip "github.com/CHIPAsia/chip-go-sdk"
)

// CREATE .env and place it inside examples folder
// Add these lines:
//
//	 	AccessKey=<<YOUR_ACCESS_KEY>>
//		BrandID=<<YOUR_BRAND_ID>>
var apiClient = chip.NewAPIClient(os.Getenv("AccessKey"), os.Getenv("BrandID"), "")

const (
	port = ":8080"
)

func main() {
	http.HandleFunc("/payment-methods", func(w http.ResponseWriter, r *http.Request) {
		result := getPaymentMethods()
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/purchases", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			id := r.URL.Query().Get("id")
			result := getPurchase(id)
			json.NewEncoder(w).Encode(result)
		}
	})

	log.Printf("Run server: localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func getPaymentMethods() chip.PaymentMethod {
	options := chip.PaymentMethodOption{
		Currency: "MYR",
	}
	result, _ := apiClient.PaymentMethods().GetMethods(options)
	return result

}

func getPurchase(id string) chip.Purchase {
	result, _ := apiClient.Purchases().GetPurchase(id)
	return result
}
