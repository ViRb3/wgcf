package resp

import "time"

type UpdateLicenseData struct {
	ID                       string    `json:"id"`
	AccountType              string    `json:"account_type"`
	Created                  time.Time `json:"created"`
	Updated                  time.Time `json:"updated"`
	PremiumData              int64     `json:"premium_data"`
	Quota                    int64     `json:"quota"`
	WarpPlus                 bool      `json:"warp_plus"`
	ReferralCount            int64     `json:"referral_count"`
	ReferralRenewalCountdown int64     `json:"referral_renewal_countdown"`
	Role                     string    `json:"role"`
}
