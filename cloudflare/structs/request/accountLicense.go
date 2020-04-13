package request

type AccountLicenseData struct {
	LicenseKey string `json:"license"` // account license key, used to bind device to specific account
}
