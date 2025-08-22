package util

import (
	"crypto/rand"
	"fmt"
	"time"
)

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

func IsHttp500Error(err error) bool {
	return err.Error() == "500 Internal Server Error"
}
