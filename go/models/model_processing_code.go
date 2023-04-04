package models

type ProcessingCode struct {
	TipoTransacao          string `json:"tipo_transacao,omitempty"`
	SourceAccountType      string `json:"source_account_type"      validate:"required"`
	DestinationAccountType string `json:"destination_account_type" validate:"required"`
}
