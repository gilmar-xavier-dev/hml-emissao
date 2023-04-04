package models

type InstallmentDetails struct {
	FinType                     string  `json:"finType"`
	FareAmount                  *Amount `json:"fare_amount,omitempty"`
	InsuranceAmount             *Amount `json:"insurance_amount,omitempty"`
	ThirdPartiesPaymntAmount    *Amount `json:"third_parties_paymnt_amount,omitempty"`
	RecordsPaymentsAmount       *Amount `json:"records_payments_amount,omitempty"`
	IssuerTotalCalculatedAmount *Amount `json:"issuer_total_calculated_amount,omitempty"`
	FirstPaymntDate             string  `json:"first_paymnt_date,omitempty"`
	InstalmntNbr                int32   `json:"instalmnt_nbr"`
	MonthlyInterestRate         int32   `json:"monthly_interest_rate,omitempty"`
	TotalEffectiveCostRateCet   int32   `json:"total_effective_cost_rate_cet,omitempty"`
	InstalmntAmount             *Amount `json:"instalmnt_amount"`
	AnnualInterestRate          int32   `json:"annual_interest_rate,omitempty"`
	InputValue                  *Amount `json:"input_value,omitempty"`
}
