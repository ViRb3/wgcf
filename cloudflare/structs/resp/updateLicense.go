package resp

import "time"

type UpdateLicenseResponse struct {
	Data UpdateLicenseData `json:"result"`
	Success bool           `json:"success"`
}

type UpdateLicenseData struct {
	ID                       string    `json:"id"`
	AccountType              string    `json:"account_type"`
	Created                  time.Time `json:"created"`
	Updated                  time.Time `json:"updated"`
	PremiumData              int       `json:"premium_data"`
	Quota                    int       `json:"quota"`
	WarpPlus                 bool      `json:"warp_plus"`
	ReferralCount            int       `json:"referral_count"`
	ReferralRenewalCountdown int       `json:"referral_renewal_countdown"`
	Role                     string    `json:"role"`
}
