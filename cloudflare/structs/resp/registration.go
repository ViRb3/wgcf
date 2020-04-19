package resp

type RegistrationData struct {
	DeviceData
	RegistrationAccount `json:"account"`
	AccessToken         string `json:"token"`
}

type RegistrationAccount struct {
	AccountData
	Usage int64 `json:"usage"`
}
