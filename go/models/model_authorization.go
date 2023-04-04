package models

type Authorization struct {
	Code        string `json:"code"        validate:"required"`
	Description string `json:"description" validate:"required"`
}
