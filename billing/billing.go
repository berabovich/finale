package billing

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
)

var filePath = "./billing.data"

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

func BillingGet() BillingData {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("cannot read file")
	}
	var n uint8
	var billing BillingData
	var d [6]bool
	for i, j := range f {
		if j == 49 {
			n += uint8(math.Pow(2, float64(len(f)-1-i)))
		}
		d[i], _ = strconv.ParseBool(string(j))
	}
	billing.CreateCustomer = d[0]
	billing.Purchase = d[1]
	billing.Payout = d[2]
	billing.Recurring = d[3]
	billing.FraudControl = d[4]
	billing.CheckoutPage = d[5]

	return billing
}
