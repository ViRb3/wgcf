package resp

import (
	"github.com/ViRb3/wgcf/cloudflare/structs/request"
)

type AccountData struct {
	request.AccountLicenseData
	UpdateLicenseData
}
