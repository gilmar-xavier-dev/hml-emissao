package models

type Chargeback struct {
	ChargebackId            string                  `json:"chargeback_id"             validate:"required"`
	OriginalPurchaseId      string                  `json:"original_purchase_id"      validate:"required"`
	AccountId               string                  `json:"account_id"                validate:"required"`
	PsProductCode           string                  `json:"psProductCode"             validate:"required"`
	PsProductName           string                  `json:"psProductName,omitempty"`
	ProductType             string                  `json:"productType,omitempty"`
	CountryCode             string                  `json:"countryCode"               validate:"required"`
	Source                  *Source                 `json:"source"                    validate:"required"`
	CallingSystemName       string                  `json:"callingSystemName"         validate:"required"`
	OriginalAuthorizationId int32                   `json:"original_authorization_id" validate:"required"`
	Authorization           *Authorization          `json:"authorization"             validate:"required"`
	Brand                   *Brand                  `json:"brand"                     validate:"required"`
	Card                    *Card                   `json:"card"                      validate:"required"`
	ChargebackMode          string                  `json:"chargeback_mode"           validate:"required"`
	OriginalAmount          *Amount                 `json:"original_amount"           validate:"required"`
	ChargebackAmount        *Amount                 `json:"chargeback_amount"         validate:"required"`
	EntryMode               *EntryMode              `json:"entry_mode"                validate:"required"`
	ChargebackReason        *CancellationReason     `json:"chargeback_reason"         validate:"required"`
	Iso8583Message          *Iso8583Message         `json:"iso8583_message"           validate:"required"`
	Nrid                    string                  `json:"nrid,omitempty"`
	ForceSynchronous        bool                    `json:"forceSynchronous,omitempty"`
	AdditionalTerminalData  *AdditionalTerminalData `json:"additionalTerminalData,omitempty"`
	TokenPaymentData        *TokenPaymentData       `json:"tokenPaymentData,omitempty"`
	HceTransaction          bool                    `json:"hceTransaction,omitempty"`
}
