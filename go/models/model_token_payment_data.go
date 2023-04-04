package models

// Informações dos Dados Transacionais - Conjunto de dados de Pagamento via Token.
type TokenPaymentData struct {
	FPan                                         string `json:"fPan,omitempty"`
	RequesterIdToken                             string `json:"requesterIdToken,omitempty"`
	Pan                                          string `json:"pan,omitempty"`
	TokenPSN                                     string `json:"tokenPSN,omitempty"`
	TokenExpirationDate                          string `json:"tokenExpirationDate,omitempty"`
	TokenStatus                                  string `json:"tokenStatus,omitempty"`
	TokenCryptogramVerificationResults           string `json:"tokenCryptogramVerificationResults,omitempty"`
	EMVTokenCryptogramVerificationResults        string `json:"EMVTokenCryptogramVerificationResults,omitempty"`
	TokenConstraintsVerificationStatus           string `json:"tokenConstraintsVerificationStatus,omitempty"`
	TransactionDateTimeConstraint                string `json:"transactionDateTimeConstraint,omitempty"`
	TransactionAmountConstraint                  string `json:"transactionAmountConstraint,omitempty"`
	UsageConstraint                              string `json:"usageConstraint,omitempty"`
	TokenATCVerificationResults                  string `json:"tokenATCVerificationResults,omitempty"`
	CVE2TokenCryptogramVerificationStatus        string `json:"CVE2TokenCryptogramVerificationStatus,omitempty"`
	MerchantVerification                         string `json:"merchantVerification,omitempty"`
	MagstripeTokenCryptogramVerificationResults  string `json:"magstripeTokenCryptogramVerificationResults,omitempty"`
	CVE2OutputTokenCryptogramVerificationResults string `json:"CVE2OutputTokenCryptogramVerificationResults,omitempty"`
}
