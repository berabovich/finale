package data

var countryAlfa = map[string]bool{
	"RU": true,
	"US": true,
	"GB": true,
	"FR": true,
	"BL": true,
	"AT": true,
	"BG": true,
	"DK": true,
	"CA": true,
	"ES": true,
	"CH": true,
	"TR": true,
	"PE": true,
	"NZ": true,
	"MC": true,
}
var countryFull = map[string]string{
	"RU": "Russia",
	"US": "USA",
	"GB": "Great Britain",
	"FR": "France",
	"BL": "Saint Bartholemy",
	"AT": "Austria",
	"BG": "Bulgaria",
	"DK": "Denmark",
	"CA": "Canada",
	"ES": "Spain",
	"CH": "Switzerland",
	"TR": "Turkey",
	"PE": "Peru",
	"NZ": "New Zealand",
	"MC": "Monaco",
}

var providers = map[string]bool{
	"Topolo": true,
	"Rond":   true,
	"Kildy":  true,
}
var providersVoice = map[string]bool{
	"TransparentCalls": true,
	"E-Voice":          true,
	"JustPhone":        true,
}
var providersEmail = map[string]bool{
	"Gmail":      true,
	"Yahoo":      true,
	"Hotmail":    true,
	"MSN":        true,
	"Orange":     true,
	"Comcast":    true,
	"AOL":        true,
	"Live":       true,
	"RediffMail": true,
	"GMX":        true,
	"Protonmail": true,
	"Yandex":     true,
	"Mail.ru":    true,
}

func CountryCheck(alfa string) bool {
	if countryAlfa[alfa] {
		return true
	}
	return false
}
func CountryAlphaToFull(alfa string) string {
	country := countryFull[alfa]
	return country
}
func ProvidersSmsMmsCheck(provider string) bool {
	if providers[provider] {
		return true
	}
	return false
}
func ProvidersVoiceCheck(provider string) bool {
	if providersVoice[provider] {
		return true
	}
	return false
}
func ProvidersEmailCheck(provider string) bool {
	if providersEmail[provider] {
		return true
	}
	return false
}
