package voice_call

import (
	"Finale/data"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var filePath = "C:/Users/k.semerenko/GolandProjects/simulator/skillbox-diploma/voice.data"

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}

func (s *VoiceCallData) parse(in string) bool {
	tmp := strings.Split(in, ";")
	if len(tmp) == 8 {
		s.Country = tmp[0]
		s.Bandwidth = tmp[1]
		s.ResponseTime = tmp[2]
		s.Provider = tmp[3]
		connectionStability, err := strconv.ParseFloat(tmp[4], 32)
		if err != nil {
			return false
		}
		s.ConnectionStability = float32(connectionStability)
		s.TTFB, err = strconv.Atoi(tmp[5])
		if err != nil {
			return false
		}
		s.VoicePurity, err = strconv.Atoi(tmp[6])
		if err != nil {
			return false
		}
		s.MedianOfCallsTime, err = strconv.Atoi(tmp[7])
		if err != nil {
			return false
		}

		return true
	}
	return false
}

func VoiceCallGet() []VoiceCallData {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("cannot read file")
	}
	t := strings.Split(string(f), "\n")
	var voiceData []VoiceCallData
	for _, d := range t {

		var voice VoiceCallData
		b := voice.parse(d)
		if b != true {
			continue
		}
		if data.CountryCheck(voice.Country) == false {
			continue
		}
		bandWidth, err := strconv.Atoi(voice.Bandwidth)
		if err != nil {
			continue
		}
		if bandWidth < 0 || bandWidth > 100 {
			continue
		}
		responseTime, err := strconv.Atoi(voice.ResponseTime)
		if err != nil {
			continue
		}
		if responseTime < 0 {
			continue
		}
		if data.ProvidersVoiceCheck(voice.Provider) == false {
			continue
		}
		voiceData = append(voiceData, voice)
	}
	return voiceData
}
