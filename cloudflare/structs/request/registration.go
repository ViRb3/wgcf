package request

type RegistrationData struct {
	PublicKey string `json:"key"` // WireGuard public key
	InstallID string `json:"install_id"`
	FcmToken  string `json:"fcm_token"` // Google Firebase token
	Tos       string `json:"tos"`       // RFC3339Nano timestamp of date accepted
	Model     string `json:"model"`
	Type      string `json:"type"`
	Locale    string `json:"locale"`
}
