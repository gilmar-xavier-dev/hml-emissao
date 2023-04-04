package models

type Establishment struct {
	Mcc     string `json:"mcc" validate:"required"`
	Name    string `json:"name,omitempty"`
	City    string `json:"city,omitempty"`
	Address string `json:"address,omitempty"`
	Zipcode string `json:"zipcode,omitempty"`
	Country string `json:"country,omitempty"`
	Cnpj    string `json:"cnpj,omitempty"`
}
