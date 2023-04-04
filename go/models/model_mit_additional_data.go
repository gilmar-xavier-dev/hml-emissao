package models

type MitAdditionalData struct {
	InstallmentTotalNbr  string `json:"installmentTotalNbr,omitempty"`
	PaymentType          string `json:"paymentType,omitempty"`
	TransactionType      string `json:"transactionType,omitempty"`
	TransactionAmountMIT int64  `json:"transactionAmountMIT,omitempty"`
	UniqueTransactionID  string `json:"uniqueTransactionID,omitempty"`
	TransactionFrequency string `json:"transactionFrequency,omitempty"`
	ValidationIndicator  string `json:"validationIndicator,omitempty"`
	ValidationReference  string `json:"validationReference,omitempty"`
	SequenceIndicator    int32  `json:"sequenceIndicator,omitempty"`
}
