package email

import (
	"Finale/internal/data"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var filePath = "../simulator/email.data"

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

func (s *EmailData) parse(in string) bool {
	tmp := strings.Split(in, ";")
	if len(tmp) == 3 {
		s.Country = tmp[0]
		s.Provider = tmp[1]
		deliveryTime, err := strconv.Atoi(tmp[2])
		if err != nil {
			return false
		}
		s.DeliveryTime = deliveryTime

		return true
	}
	return false
}

// EmailGet get email data from file and check it
func EmailGet() []EmailData {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("cannot read file")
	}
	t := strings.Split(string(f), "\n")
	var emailData []EmailData
	for _, d := range t {
		var email EmailData
		b := email.parse(d)
		if b != true {
			continue
		}
		if data.CountryCheck(email.Country) == false {
			continue
		}
		if data.ProvidersEmailCheck(email.Provider) == false {
			continue
		}
		emailData = append(emailData, email)
	}
	return emailData
}
