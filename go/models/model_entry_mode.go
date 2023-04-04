package models

type EntryMode string

const (
	UNKNOWN         EntryMode = "unknown"
	MAGNETIC_STRIPE EntryMode = "magnetic_stripe"
	FALLBACK        EntryMode = "fallback"
	CHIP            EntryMode = "chip"
	ECOMMERCE       EntryMode = "ecommerce"
	MCOMMERCE       EntryMode = "mcommerce"
	TYPED           EntryMode = "typed"
	CONTACTLESS     EntryMode = "contactless"
	STORED_ACCOUNT  EntryMode = "stored_account"
	BAR_OR_QR_CODE  EntryMode = "bar_or_qr_code"
	URA             EntryMode = "ura"
)
