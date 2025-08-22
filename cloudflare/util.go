package cloudflare

import (
	"github.com/cockroachdb/errors"
)

func FindDevice(devices []BoundDevice, deviceId string) (*BoundDevice, error) {
	for _, device := range devices {
		if device.Id == deviceId {
			return &device, nil
		}
	}
	return nil, errors.New("device not found in list")
}
