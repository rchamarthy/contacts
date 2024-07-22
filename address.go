package contacts

type Address struct {
	AptNumber  string `json:"aptNumber" yaml:"aptNumber"`
	Street     string `json:"street" yaml:"street"`
	City       string `json:"city" yaml:"city"`
	State      string `json:"state" yaml:"state"`
	PostalCode string `json:"postalCode" yaml:"postalCode"`
	Country    string `json:"country" yaml:"country"`
}
