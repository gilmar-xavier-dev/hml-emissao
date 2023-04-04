package models

type Card struct {
	PaysmartId string `json:"paysmart_id" validate:"required"`
	IssuerId   string `json:"issuer_id,omitempty"`
	Pan        string `json:"pan"         validate:"required"`
	Panseq     string `json:"panseq,omitempty"`
}
