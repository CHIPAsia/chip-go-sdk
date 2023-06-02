# Chip Golang library

## Prerequisite
Before you start, make sure you already have created `Brand ID` and `API Key` from your developer dashboard by logging-in into [merchant portal](https://gate.chip-in.asia/login).

## Installation
```golang
$ go get -u github.com/CHIPAsia/chip-go-sdk
```

## Getting Started
Simple usage:

```golang
import chip "github.com/CHIPAsia/chip-go-sdk"

func main() {
  apiKey := "<<YOUR_API_KEY>>"
  brandID := "<<YOUR_BRAND_ID>>"
  apiClient := chip.NewAPIClient(apiKey, brandID, "")
  
  client := chip.Client{
    Email:    "test@test.com",
    Fullname: "John Doe",
    Phone:    "60111222333",
  }

  purchase := chip.PurchaseDetail{
    products: []chip.Product{
      {
        Name: "Product 1",
        Quantity: 1,
        Price: 1000, // 10.00
      }
    }
  }

  purchase := PurchaseOption {
		Client: client,
		Purchase: purchase,
		SuccessRedirect: "https://example.com/redirect?status=success",
		FailureRedirect: "https://example.com/redirect?status=fail",
		CancelRedirect: "https://example.com/redirect?status=cancel",
		SuccessCallback: "https://example.com/callback",
	}

  if response, err := apiClient.Purchases().CreatePurchase(purchase); err == nil {
    fmt.Printf("Response %+v\n", response)
  }
}
```

## Example
Checkout our examples [here](./examples).

To run our example:
1. Create file with name `.env` and add these two lines:
```bash
AccessKey=<<YOUR_API_KEY>>
BrandID=<<YOUR_BRAND_ID>>
```
2. Run `./scripts/run-example.sh` from your terminal.