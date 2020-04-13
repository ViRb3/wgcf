package url

import . "wgcf/util"

var apiVersion = "v0a884"
var apiEndpoint = JoinUrls("https://api.cloudflareclient.com", apiVersion)

var RegUrl = JoinUrls(apiEndpoint, "reg")
var ClientConfigUrl = JoinUrls(apiEndpoint, "client_config")

var TermsUrl = "https://www.cloudflare.com/application/terms/"

// returns DeviceData
func GetDeviceUrl(deviceId string) string {
	return JoinUrls(RegUrl, deviceId)
}

// returns AccountData
func GetAccountUrl(deviceId string) string {
	return JoinUrls(GetDeviceUrl(deviceId), "account")
}

// returns BoundDevicesData
func GetBoundDevicesUrl(deviceId string) string {
	return JoinUrls(GetAccountUrl(deviceId), "devices")
}

// creates and returns a new license key that replaces the old key
func GetRecreateLicenseKeyUrl(deviceId string) string {
	return JoinUrls(GetAccountUrl(deviceId), "license")
}

// allows operating on devices bound to this device's account
func GetBoundDeviceUrl(deviceId string, targetDeviceId string) string {
	return JoinUrls(GetAccountUrl(deviceId), "reg", targetDeviceId)
}
