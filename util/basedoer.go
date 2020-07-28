package util

import (
	"github.com/dghubble/sling"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type BaseDoer struct {
	*http.Client
}

func (b *BaseDoer) Do(req *http.Request) (*http.Response, error) {
	response, err := b.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		return nil, CreateNon200Error(response.StatusCode, bodyBytes)
	}
	return response, nil
}

func newBaseClient() *http.Client {
	return &http.Client{Timeout: time.Duration(5000) * time.Second}
}

func NewBaseDoer() (sling.Doer, *http.Client) {
	client := newBaseClient()
	doer := BaseDoer{Client: client}
	return &doer, client
}

func NewJarDoer() (sling.Doer, *http.Client) {
	jar, _ := cookiejar.New(nil)
	client := newBaseClient()
	client.Jar = jar
	doer := BaseDoer{Client: client}
	return &doer, client
}
