package resp

type ConfigData struct {
	PremiumDataBytes    int64      `json:"premium_data_bytes"`
	ReferralRewardBytes int64      `json:"referral_reward_bytes"`
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
