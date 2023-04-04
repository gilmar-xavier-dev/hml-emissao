package models

type Amount struct {
	Amount       int64 `json:"amount"        validate:"required"`
	CurrencyCode int32 `json:"currency_code" validate:"required"`
}
