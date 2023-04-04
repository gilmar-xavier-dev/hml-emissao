package models

type HolderValidationMode string

const (
	ONLINE_PIN  HolderValidationMode = "online_pin"
	OFFLINE_PIN HolderValidationMode = "offline_pin"
	OTHER       HolderValidationMode = "other"
)
