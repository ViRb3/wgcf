package resp

import (
	"time"
	"wgcf/cloudflare/structs/request"
)

type DeviceData struct {
	request.RegistrationData
	request.DeviceWarpData
	request.DeviceNameData
	ID              string      `json:"id"`
	Account         AccountData `json:"account"`
	Config          Config      `json:"config"`
	WaitlistEnabled bool        `json:"waitlist_enabled"`
	Created         time.Time   `json:"created"`
	Updated         time.Time   `json:"updated"`
	Place           int64       `json:"place"`
	Enabled         bool        `json:"enabled"`
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
