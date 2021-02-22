package util

import (
	"crypto/rand"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
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

// Casts a type into another type.
// Works like mapstructure, but more reliable.
func Restructure(source interface{}, dest interface{}) error {
	bytes, err := yaml.Marshal(source)
	if err != nil {
		return errors.WithMessage(err, "marshal")
	}
	if err := yaml.Unmarshal(bytes, dest); err != nil {
		return errors.WithMessage(err, "unmarshal")
	}
	return nil
}
