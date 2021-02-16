package efi

// EnvironmentType represents type of Environment
type EnvironmentType string

// EnvironmentType constants
const (
	TEST EnvironmentType = "TEST"
	PROD EnvironmentType = "PROD"
)

// Config represents configuration that is used by efi modules
type Config struct {
	// Company details
	TIN         string `json:"TIN"`
	Name        string `json:"Name"`
	VAT         string `json:"VAT"`
	Address     string `json:"Address"`
	Town        string `json:"Town"`
	Country     string `json:"Country"`
	Phone       string `json:"Phone"`
	Fax         string `json:"Fax"`
	BankAccount string `json:"BankAccount"`
	// Environment, eg staging or production
	Environment EnvironmentType `json:"Environment"`
	// EFI-related constants
	BusinUnitCode string `json:"BusinUnitCode"`
	TCRCode       string `json:"TCRCode"`
	SoftCode      string `json:"SoftCode"`
	OperatorCode  string `json:"OperatorCode"`
	// Typless constnts
	Typless struct {
		APIKey   string `json:"APIKey"`
		Template string `json:"Template"`
	} `json:"Typless"`
}
