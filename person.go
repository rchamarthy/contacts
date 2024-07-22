package contacts

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Person captures a person's contact information
type Person struct {
	Email           string        `json:"email" yaml:"email"`
	FirstName       string        `json:"firstName" yaml:"firstName"`
	LastName        string        `json:"lastName" yaml:"lastName"`
	MiddleName      string        `json:"middleName" yaml:"middleName"`
	PhoneNumbers    []PhoneNumber `json:"phoneNumbers" yaml:"phoneNumbers"`
	Addresses       []Address     `json:"addresses" yaml:"addresses"`
	BirthDate       Date          `json:"birthDate" yaml:"birthDate"`
	NickName        string        `json:"nickName" yaml:"nickName"`
	Notes           string        `json:"notes" yaml:"notes"`
	Groups          []string      `json:"groups" yaml:"groups"`
	Occupation      string        `json:"occupation" yaml:"occupation"`
	ProfileImageURL string        `json:"profileImageURL" yaml:"profileImageURL"`
}

func EmptyPerson() *Person {
	return &Person{}
}

func (p Person) Validate() error {
	return nil
}

func LoadPerson(filename string) (*Person, error) {
	b, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	p := EmptyPerson()
	err = yaml.Unmarshal(b, &p)
	if err != nil {
		return nil, err
	}

	if e := p.Validate(); e != nil {
		return nil, e
	}

	return p, nil
}

func (p Person) Save(filename string) error {
	b, err := yaml.Marshal(&p)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, b, 0600)
}
