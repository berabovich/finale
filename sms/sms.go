package sms

import (
	"Finale/data"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var filePath = "C:/Users/k.semerenko/GolandProjects/simulator/skillbox-diploma/sms.data"

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func (s *SMSData) parse(in string) bool {
	tmp := strings.Split(in, ";")
	if len(tmp) == 4 {
		s.Country = tmp[0]
		s.Bandwidth = tmp[1]
		s.ResponseTime = tmp[2]
		s.Provider = tmp[3]
		return true
	}
	return false
}

func SmsGet() []SMSData {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("cannot read file")
	}
	t := strings.Split(string(f), "\n")
	var smsData []SMSData
	for _, d := range t {

		var sms SMSData
		b := sms.parse(d)
		if b != true {
			continue
		}
		if data.CountryCheck(sms.Country) == false {
			continue
		}
		sms.Country = data.CountryAlphaToFull(sms.Country)
		bandWidth, err := strconv.Atoi(sms.Bandwidth)
		if err != nil {
			continue
		}
		if bandWidth < 0 || bandWidth > 100 {
			continue
		}
		responseTime, err := strconv.Atoi(sms.ResponseTime)
		if err != nil {
			continue
		}
		if responseTime < 0 {
			continue
		}
		if data.ProvidersSmsMmsCheck(sms.Provider) == false {
			continue
		}
		smsData = append(smsData, sms)
	}
	return smsData
}
