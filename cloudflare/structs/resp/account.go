package resp

import (
	"wgcf/cloudflare/structs/request"
)

type AccountData struct {
	request.AccountLicenseData
	UpdateLicenseData
}
