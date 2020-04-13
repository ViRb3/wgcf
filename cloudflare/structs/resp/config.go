package resp

type ConfigData struct {
	PremiumDataBytes    int        `json:"premium_data_bytes"`
	ReferralRewardBytes int        `json:"referral_reward_bytes"`
	Denylist            []Denylist `json:"denylist"`
}
type V4 struct {
	Address string `json:"address"`
	Netmask string `json:"netmask"`
}
type Networks struct {
	V4 []V4          `json:"v4"`
	V6 []interface{} `json:"v6"`
}
type Denylist struct {
	Name            string   `json:"name"`
	AndroidPackages []string `json:"android-packages,omitempty"`
	Visible         bool     `json:"visible"`
	Networks        Networks `json:"networks,omitempty"`
}
