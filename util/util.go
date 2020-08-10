package util

import (
	"crypto/rand"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrNon200StatusCode = errors.New("non-200 status code")
)

func CreateNon200Error(code int, body []byte) error {
	return errors.WithMessage(errors.WithStack(ErrNon200StatusCode),
		fmt.Sprintf("code: %d, body: %s", code, string(body)))
}

func GetErrorMessage(err error) string {
	return fmt.Sprintf("%+v", err)
}

func RandomHexString(count int) string {
	b := make([]byte, count)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%X", b)
}

func GetTimestamp() string {
	return getTimestamp(time.Now())
}

func getTimestamp(t time.Time) string {
	timestamp := t.Format(time.RFC3339Nano)
	return timestamp
}
