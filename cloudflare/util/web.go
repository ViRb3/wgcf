package util

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
)

var defaultHeaders = map[string]string{"User-Agent": "okhttp/3.12.1"}

func applyDefaultHeaders(request *http.Request) {
	for key, value := range defaultHeaders {
		request.Header.Set(key, value)
	}
}

var client = http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			// Match app's TLS config or API will reject us with code 403 error 1020
			MinVersion: tls.VersionTLS10,
			MaxVersion: tls.VersionTLS12}}}

func NewAuthenticatedRequest(method string, url string, body io.Reader, accessToken string, result interface{}) error {
	request, err := newRequest(method, url, body, accessToken)
	if err != nil {
		return err
	}
	return doRequest(request, result)
}

func NewUnauthenticatedRequest(method string, url string, body io.Reader, result interface{}) error {
	request, err := newRequest(method, url, body, "")
	if err != nil {
		return err
	}
	return doRequest(request, result)
}

func doRequest(request *http.Request, result interface{}) error {
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 {
		return errors.WithMessage(errors.New("non-200 status code"),
			fmt.Sprintf("code: %d, body: %s", response.StatusCode, string(bodyBytes)))
	}
	return json.Unmarshal(bodyBytes, result)
}

func newRequest(method string, url string, body io.Reader, accessToken string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	applyDefaultHeaders(request)
	if body != nil {
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	if accessToken != "" {
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	}
	return request, nil
}
