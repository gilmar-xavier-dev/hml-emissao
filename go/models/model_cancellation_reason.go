package models

type CancellationReason struct {
	Code        int32  `json:"code"         validate:"required"`
	Description string `json:"description"  validate:"required"`
}
