package models

type AdditionalTerminalData struct {
	TerminalType                   string `json:"terminalType,omitempty"`
	PartialApprovalIndicator       string `json:"partialApprovalIndicator,omitempty"`
	TerminalLocationIndicator      string `json:"terminalLocationIndicator,omitempty"`
	CardholderPresenceIndicator    string `json:"cardholderPresenceIndicator,omitempty"`
	CardPresenceIndicator          string `json:"cardPresenceIndicator,omitempty"`
	CardCaptureCapabilityIndicator string `json:"cardCaptureCapabilityIndicator,omitempty"`
	TransactionStatusIndicator     string `json:"transactionStatusIndicator,omitempty"`
	TransactionSecurityIndicator   string `json:"transactionSecurityIndicator,omitempty"`
	TerminalPOSType                string `json:"terminalPOSType,omitempty"`
	TerminalInputCapability        string `json:"terminalInputCapability,omitempty"`
}
