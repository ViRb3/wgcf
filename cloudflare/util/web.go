package util

import (
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

var client = http.Client{}

func init() {
	//proxyUrl, _ := url.Parse("http://127.0.0.1:8888")
	//client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
}

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
	if response.StatusCode != 200 {
		return errors.New(fmt.Sprintf("non-200 status code: %d", response.StatusCode))
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
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
