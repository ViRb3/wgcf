package api

import (
	"bytes"
	"encoding/json"
	"wgcf/cloudflare/structs/request"
	"wgcf/cloudflare/structs/resp"
	"wgcf/cloudflare/url"
	"wgcf/cloudflare/util"
	"wgcf/config"
)

func UpdateLicenseKey(ctx *config.Context) (*resp.UpdateLicenseData, error) {
	data := request.AccountLicenseData{LicenseKey: ctx.LicenseKey}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var result resp.UpdateLicenseData
	if err := util.NewAuthenticatedRequest("PUT", url.GetAccountUrl(ctx.DeviceId),
		bytes.NewBuffer(dataBytes), ctx.AccessToken, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// returns the BoundDevice entry for a device, used to check additional settings such as name and active status
func GetBoundDevice(ctx *config.Context) (*resp.BoundDevice, error) {
	var result resp.BoundDevicesData
	if err := util.NewAuthenticatedRequest("GET", url.GetBoundDevicesUrl(ctx.DeviceId), nil,
		ctx.AccessToken, &result); err != nil {
		return nil, err
	}
	return util.FindDevice(&result, ctx.DeviceId)
}

// only 5 linked devices can be active at a given time
func SetBoundDeviceActive(ctx *config.Context, data *request.DeviceActiveData) (*resp.BoundDevice, error) {
	return setBoundDeviceData(ctx, data)
}

func SetBoundDeviceName(ctx *config.Context, data *request.DeviceNameData) (*resp.BoundDevice, error) {
	return setBoundDeviceData(ctx, data)
}

func setBoundDeviceData(ctx *config.Context, data interface{}) (*resp.BoundDevice, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var result resp.BoundDevicesData
	if err := util.NewAuthenticatedRequest("PATCH", url.GetBoundDeviceUrl(ctx.DeviceId, ctx.DeviceId),
		bytes.NewBuffer(dataBytes), ctx.AccessToken, &result); err != nil {
		return nil, err
	}
	device, err := util.FindDevice(&result, ctx.DeviceId)
	if err != nil {
		return nil, err
	}
	return device, nil
}
