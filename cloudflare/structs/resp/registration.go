package resp

type RegistrationResponse struct {
	Data RegistrationData `json:"result"`
	Success bool          `json:"success"`
}

type RegistrationData struct {
	DeviceData
	RegistrationAccount `json:"account"`
	AccessToken         string `json:"token"`
}

type RegistrationAccount struct {
	AccountData
	Usage int `json:"usage"`
}
