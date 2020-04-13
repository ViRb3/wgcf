package resp

import (
	"time"
	"wgcf/cloudflare/structs/request"
)

type BoundDevicesData []BoundDevice

type BoundDevice struct {
	request.DeviceActiveData
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Model     string    `json:"model"`
	Name      string    `json:"name,omitempty"`
	Created   time.Time `json:"created"`
	Activated time.Time `json:"activated"`
	Active    bool      `json:"active"`
	Role      string    `json:"role"`
}
