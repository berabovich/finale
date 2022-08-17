package result

import (
	"Finale/internal/billing"
	"Finale/internal/data"
	"Finale/internal/email"
	"Finale/internal/incident"
	"Finale/internal/mms"
	"Finale/internal/sms"
	"Finale/internal/support"
	"Finale/internal/voice_call"
	"sort"
)

type ResultT struct {
	Status bool `json:"status"` // true, если все этапы сбора данных прошли успешно, false во всех остальных случаях

	Data ResultSetT `json:"data"` // заполнен, если все этапы сбора данных прошли успешно, nil во всех остальных случаях

	Error string `json:"error"` // пустая строка если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки (детали ниже)

}

type ResultSetT struct {
	SMS       [][]sms.SMSData                `json:"sms"`
	MMS       [][]mms.MmsData                `json:"mms"`
	VoiceCall []voice_call.VoiceCallData     `json:"voice_call"`
	Email     map[string][][]email.EmailData `json:"email"`
	Billing   billing.BillingData            `json:"billing"`
	Support   []int                          `json:"support"`
	Incidents []incident.IncidentData        `json:"incident"`
}

func GetResultData() ResultSetT {
	var r ResultSetT
	resultSMS(&r)
	resultMMS(&r)
	resultVoiceCall(&r)
	resultEmail(&r)
	resultBilling(&r)
	resultSupport(&r)
	resultIncident(&r)
	return r
}

func resultSMS(r *ResultSetT) ResultSetT {
	smsByCountry := sms.SmsGet()

	sort.Slice(smsByCountry, func(i, j int) bool {
		return smsByCountry[i].Country < smsByCountry[j].Country
	})

	smsByProvider := sms.SmsGet()

	sort.Slice(smsByProvider, func(i, j int) bool {
		return smsByProvider[i].Provider < smsByProvider[j].Provider
	})

	r.SMS = append(r.SMS, smsByProvider)
	r.SMS = append(r.SMS, smsByCountry)
	return *r
}
func resultMMS(r *ResultSetT) ResultSetT {
	mmsByCountry := mms.MMSget()

	sort.Slice(mmsByCountry, func(i, j int) bool {
		return mmsByCountry[i].Country < mmsByCountry[j].Country
	})

	mmsByProvider := mms.MMSget()

	sort.Slice(mmsByProvider, func(i, j int) bool {
		return mmsByProvider[i].Provider < mmsByProvider[j].Provider
	})

	r.MMS = append(r.MMS, mmsByProvider)
	r.MMS = append(r.MMS, mmsByCountry)
	return *r
}
func resultVoiceCall(r *ResultSetT) ResultSetT {
	r.VoiceCall = voice_call.VoiceCallGet()
	return *r
}
func resultEmail(r *ResultSetT) ResultSetT {
	emailByTime := email.EmailGet()
	sor := make(map[string][]email.EmailData)
	for _, s := range emailByTime {
		sor[s.Country] = append(sor[s.Country], s)
	}
	e := make(map[string][][]email.EmailData)
	for t, s := range sor {
		sort.Slice(s, func(i, j int) bool {
			return s[i].DeliveryTime < s[j].DeliveryTime
		})

		high := s[:3]
		low := s[len(s)-3:]
		var g [][]email.EmailData
		g = append(g, high)
		g = append(g, low)
		e[data.CountryAlphaToFull(t)] = g

	}
	r.Email = e
	return *r
}
func resultBilling(r *ResultSetT) ResultSetT {
	r.Billing = billing.BillingGet()
	return *r
}
func resultSupport(r *ResultSetT) ResultSetT {
	supportData := support.SupportGet()
	load := 0
	sup := 0
	for _, i := range supportData {
		sup += i.ActiveTickets
	}
	switch {
	case sup < 9:
		load = 1
	case sup > 9 && sup < 16:
		load = 2
	case sup > 16:
		load = 3
	}
	time := 60 / 18 * sup
	r.Support = append(r.Support, load)
	r.Support = append(r.Support, time)
	return *r
}
func resultIncident(r *ResultSetT) ResultSetT {
	incidentData := incident.IncidentGet()
	sort.Slice(incidentData, func(i, j int) bool {
		return incidentData[i].Status < incidentData[j].Status
	})
	r.Incidents = incidentData
	return *r
}
