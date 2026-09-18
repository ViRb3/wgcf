package wireguard

import (
	"encoding/base64"
	"testing"
)

func TestNewKey(t *testing.T) {
	key, err := NewPrivateKey()
	if err != nil {
		t.Error(err)
	}
	encodedKey := key.String()
	newKey, err := NewKey(encodedKey)
	if err != nil {
		t.Error(err)
	}
	if newKey.String() != encodedKey {
		t.Error()
	}
}

func TestNewKeyRejectsWrongLength(t *testing.T) {
	shortKey := base64.StdEncoding.EncodeToString(make([]byte, KeyLength-1))
	if _, err := NewKey(shortKey); err == nil {
		t.Fatal("NewKey accepted a 31-byte key")
	}
	if IsKey(shortKey) {
		t.Fatal("IsKey accepted a 31-byte key")
	}
}
