package cloudflare

type RegRequest struct {
	PublicKey string `json:"key"`
	InstallID string `json:"install_id"`
	FcmToken  string `json:"fcm_token"`
	Tos       string `json:"tos"`
	Model     string `json:"model"`
	Type      string `json:"type"`
	Locale    string `json:"locale"`
}

type UpdateLicenseKeyRequest struct {
	LicenseKey string `json:"license"`
}

type SetBoundDeviceNameRequest struct {
	Name string `json:"name"`
}

type SetBoundDeviceActiveRequest struct {
	Active bool `json:"active"`
}
type UpdateLicenseKeyRequest2 struct {
	PublicKey string `json:"key"`
}
