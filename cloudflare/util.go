package cloudflare

import (
	"github.com/pkg/errors"
)

func FindDevice(devices []BoundDevice, deviceId string) (*BoundDevice, error) {
	for _, device := range devices {
		if device.Id == deviceId {
			return &device, nil
		}
	}
	return nil, errors.New("device not found in list")
}
