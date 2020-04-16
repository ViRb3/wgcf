package util

import (
	"fmt"
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

func GetErrorMessage(err error) string {
	return fmt.Sprintf("%+v", err)
}
