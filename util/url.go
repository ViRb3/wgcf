package util

import (
	"path"
	"strings"
)

func JoinUrls(baseUrl string, extensionUrl ...string) string {
	extensionUrlFull := path.Join(extensionUrl...)
	requestUrl := path.Join(baseUrl, extensionUrlFull)
	// fix https:/website.com -> https://website.com
	requestUrl = strings.Replace(requestUrl, "/", "//", 1)
	return requestUrl
}
