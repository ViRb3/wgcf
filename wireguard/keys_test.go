package wireguard

import "testing"

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
