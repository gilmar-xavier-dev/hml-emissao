package models

type Purchase struct {
	PurchaseId               string                  `json:"purchase_id"              validate:"required"`
	AccountId                string                  `json:"account_id"               validate:"required"`
	PsProductCode            string                  `json:"psProductCode"            validate:"required"`
	PsProductName            string                  `json:"psProductName,omitempty"`
	ProductType              string                  `json:"productType,omitempty"`
	CountryCode              string                  `json:"countryCode"              validate:"required"`
	Source                   *Source                 `json:"source"                   validate:"required"`
	CallingSystemName        string                  `json:"callingSystemName"        validate:"required"`
	PreAuthorization         bool                    `json:"preAuthorization"         validate:"boolean"`
	IncrementalAuthorization bool                    `json:"incrementalAuthorization" validate:"boolean"`
	Authorization            *Authorization          `json:"authorization"            validate:"required"`
	Brand                    *Brand                  `json:"brand"                    validate:"required"`
	Card                     *Card                   `json:"card"                     validate:"required"`
	TotalAmount              *Amount                 `json:"total_amount"             validate:"required"`
	DollarAmount             *Amount                 `json:"dollar_amount,omitempty"`
	OriginalAmount           *Amount                 `json:"original_amount"          validate:"required"`
	DollarRealRate           string                  `json:"dollar_real_rate,omitempty"`
	Spread                   string                  `json:"spread,omitempty"`
	EntryMode                *EntryMode              `json:"entry_mode"               validate:"required"`
	ProcessingCode           *ProcessingCode         `json:"processing_code"          validate:"required"`
	HolderValidationMode     *HolderValidationMode   `json:"holder_validation_mode,omitempty"`
	Fees                     []Fee                   `json:"fees"                     validate:"required"`
	Establishment            *Establishment          `json:"establishment"            validate:"required"`
	Internacional            bool                    `json:"internacional"            validate:"boolean"`
	OriginalIso8583          *Iso8583Message         `json:"original_iso8583"         validate:"required"`
	ForceAccept              bool                    `json:"forceAccept,omitempty"`
	FraudData                *FraudData              `json:"fraudData,omitempty"`
	InstallmentDetails       *InstallmentDetails     `json:"installmentDetails,omitempty"`
	Nrid                     string                  `json:"nrid,omitempty"`
	AuthorizationAdvice      bool                    `json:"authorizationAdvice,omitempty"`
	MitAdditionalData        *MitAdditionalData      `json:"mitAdditionalData,omitempty"`
	AdditionalTerminalData   *AdditionalTerminalData `json:"additionalTerminalData,omitempty"`
	TokenPaymentData         *TokenPaymentData       `json:"tokenPaymentData,omitempty"`
	HceTransaction           bool                    `json:"hceTransaction,omitempty"`
}
