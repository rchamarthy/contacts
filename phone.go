package contacts

type PhoneNumber struct {
	CountryCode string `json:"countryCode" yaml:"countryCode"`
	Number      string `json:"number" yaml:"number"`
	Type        string `json:"type" yaml:"type"`
}
