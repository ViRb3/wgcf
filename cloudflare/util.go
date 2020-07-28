package cloudflare

import (
	"github.com/pkg/errors"
	"time"
)

func GetTimestamp() string {
	return getTimestamp(time.Now())
}

func getTimestamp(t time.Time) string {
	timestamp := t.Format(time.RFC3339Nano)
	return timestamp
}

func FindDevice(devices []Device, deviceId string) (*Device, error) {
	for _, device := range devices {
		if device.ID == deviceId {
			return &device, nil
		}
	}
	return nil, errors.New("device not found in list")
}
