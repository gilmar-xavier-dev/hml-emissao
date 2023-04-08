package models

type Query struct {
	QueryId                string                  `json:"query_id"                validate:"required"`
	AccountId              string                  `json:"account_id"              validate:"required"`
	PsProductCode          string                  `json:"psProductCode"           validate:"required"`
	PsProductName          string                  `json:"psProductName,omitempty"`
	ProductType            string                  `json:"productType,omitempty"`
	CountryCode            string                  `json:"countryCode"             validate:"required"`
	Source                 *Source                 `json:"source"                  validate:"required"`
	CallingSystemName      string                  `json:"callingSystemName"       validate:"required"`
	Authorization          *Authorization          `json:"authorization"           validate:"required"`
	Brand                  *Brand                  `json:"brand"                   validate:"required"`
	Card                   *Card                   `json:"card"                    validate:"required"`
	EntryMode              *EntryMode              `json:"entry_mode"              validate:"required"`
	Establishment          *Establishment          `json:"establishment"           validate:"required"`
	Internacional          bool                    `json:"internacional"           validate:"required"`
	OriginalIso8583        *Iso8583Message         `json:"original_iso8583"        validate:"required"`
	Nrid                   string                  `json:"nrid,omitempty"`
	AdditionalTerminalData *AdditionalTerminalData `json:"additionalTerminalData,omitempty"`
	TokenPaymentData       *TokenPaymentData       `json:"tokenPaymentData,omitempty"`
	HceTransaction         bool                    `json:"hceTransaction,omitempty"`
}
