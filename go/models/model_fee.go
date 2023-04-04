package models

type Fee struct {
	Amount *Amount `json:"amount" validate:"required"`
	Type_  string  `json:"type" validate:"required"`
}
