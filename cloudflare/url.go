package cloudflare

var apiEndpoint = "https://api.cloudflareclient.com/v0a977"

var regUrl = apiEndpoint + "/reg"
var clientConfigUrl = apiEndpoint + "/client_config"

var TermsUrl = "https://www.cloudflare.com/application/terms/"

func getThisDeviceUrl(deviceId string) string {
	return regUrl + "/" + deviceId
}

func getAccountUrl(deviceId string) string {
	return getThisDeviceUrl(deviceId) + "/account"
}

func getBoundDevicesUrl(deviceId string) string {
	return getAccountUrl(deviceId) + "/devices"
}

// creates and returns a new license key that replaces the old key
func getRecreateLicenseKeyUrl(deviceId string) string {
	return getAccountUrl(deviceId) + "/license"
}

// allows operating on devices bound to this device's account
func getBoundDeviceUrl(deviceId string, targetDeviceId string) string {
	return getAccountUrl(deviceId) + "/reg/" + targetDeviceId
}
