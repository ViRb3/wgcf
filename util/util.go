package util

import (
	"crypto/rand"
	"fmt"
	"time"
)

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
