package cloudflare

import "time"

type ClientConfigResponse struct {
	PremiumDataBytes    int64      `json:"premium_data_bytes"`
	ReferralRewardBytes int64      `json:"referral_reward_bytes"`
	Denylist            []Denylist `json:"denylist"`
}
type V4 struct {
	Address string `json:"address"`
	Netmask string `json:"netmask"`
}
type V6 struct {
	Address string `json:"address"`
	Prefix  int64  `json:"prefix"`
}
type Networks struct {
	V4 []V4 `json:"v4"`
	V6 []V6 `json:"v6"`
}
type Denylist struct {
	Name            string   `json:"name"`
	AndroidPackages []string `json:"android-packages,omitempty"`
	Visible         bool     `json:"visible"`
	Networks        Networks `json:"networks,omitempty"`
}

type Account struct {
	ID                       string    `json:"id"`
	AccountType              string    `json:"account_type"`
	Created                  time.Time `json:"created"`
	Updated                  time.Time `json:"updated"`
	PremiumData              int64     `json:"premium_data"`
	Quota                    int64     `json:"quota"`
	Usage                    *int      `json:"usage"`
	WarpPlus                 bool      `json:"warp_plus"`
	ReferralCount            int64     `json:"referral_count"`
	ReferralRenewalCountdown int64     `json:"referral_renewal_countdown"`
	Role                     string    `json:"role"`
	LicenseKey               *string   `json:"license"`
}
type Endpoint struct {
	V4   string `json:"v4"`
	V6   string `json:"v6"`
	Host string `json:"host"`
}
type Peers struct {
	PublicKey string   `json:"public_key"`
	Endpoint  Endpoint `json:"endpoint"`
}
type Addresses struct {
	V4 string `json:"v4"`
	V6 string `json:"v6"`
}
type Interface struct {
	Addresses Addresses `json:"addresses"`
}
type Services struct {
	HTTPProxy string `json:"http_proxy"`
}
type Config struct {
	ClientID  string    `json:"client_id"`
	Peers     []Peers   `json:"peers"`
	Interface Interface `json:"interface"`
	Services  Services  `json:"services"`
}

type Device struct {
	ID              string     `json:"id"`
	Type            string     `json:"type"`
	Model           string     `json:"model"`
	Name            *string    `json:"name,omitempty"`
	PublicKey       *string    `json:"key"`
	Account         *Account   `json:"account"`
	Config          *Config    `json:"config"`
	AccessToken     *string    `json:"token"`
	WarpEnabled     *bool      `json:"warp_enabled"`
	WaitlistEnabled *bool      `json:"waitlist_enabled"`
	Created         time.Time  `json:"created"`
	Updated         *time.Time `json:"updated"`
	Tos             *time.Time `json:"tos"`
	Activated       *time.Time `json:"activated"`
	Active          *bool      `json:"active"`
	Place           *int64     `json:"place"`
	Locale          *string    `json:"locale"`
	Enabled         *bool      `json:"enabled"`
	InstallID       *string    `json:"install_id"`
	FcmToken        *string    `json:"fcm_token"`
	Role            *string    `json:"role"`
}
