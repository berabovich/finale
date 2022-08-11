package mms

import (
	"Finale/data"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type MmsData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func MMSget() []MmsData {
	var mmsData []MmsData
	r, err := http.Get("http://127.0.0.1:8383/mms")

	if err != nil {
		return mmsData
	}
	if r.StatusCode != http.StatusOK {
		return mmsData
	}

	f, err := io.ReadAll(r.Body)
	err = json.Unmarshal(f, &mmsData)
	for i, tmp := range mmsData {
		if data.CountryCheck(tmp.Country) == false {
			mmsData = append(mmsData[:i-1], mmsData[i:]...)
			continue
		}
		mmsData[i].Country = data.CountryAlphaToFull(tmp.Country)
		if data.ProvidersSmsMmsCheck(tmp.Provider) == false {
			mmsData = append(mmsData[:i-1], mmsData[i:]...)
			continue
		}
		bandWidth, err := strconv.Atoi(tmp.Bandwidth)
		if err != nil {
			mmsData = append(mmsData[:i-1], mmsData[i:]...)
			continue
		}
		if bandWidth < 0 || bandWidth > 100 {
			continue
		}
		responseTime, err := strconv.Atoi(tmp.ResponseTime)
		if err != nil {
			mmsData = append(mmsData[:i-1], mmsData[i:]...)
			continue
		}
		if responseTime < 0 {
			mmsData = append(mmsData[:i-1], mmsData[i:]...)
			continue
		}

	}
	return mmsData
}
