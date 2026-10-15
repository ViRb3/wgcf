package wireguard

import (
	"strings"
	"testing"
)

func TestGenerateProfile(t *testing.T) {
	var expectedResult = `[Interface]
PrivateKey = 1
Address = 2/32, 3/128
DNS = 1.1.1.1, 1.0.0.1, 2606:4700:4700::1111, 2606:4700:4700::1001
MTU = 1280
[Peer]
PublicKey = 4
AllowedIPs = 0.0.0.0/0, ::/0
Endpoint = 5
`

	result, err := generateProfile(&ProfileData{
		PrivateKey: "1",
		Address1:   "2",
		Address2:   "3",
		PublicKey:  "4",
		Endpoint:   "5",
		DNS:        "1.1.1.1, 1.0.0.1, 2606:4700:4700::1111, 2606:4700:4700::1001",
	})
	if err != nil {
		t.Error(err)
	}

	if expectedResult != result {
		t.Error()
	}
}

func TestGenerateProfileWithKeepalive(t *testing.T) {
	var expectedResult = `[Interface]
PrivateKey = 1
Address = 2/32, 3/128
DNS = 1.1.1.1, 1.0.0.1, 2606:4700:4700::1111, 2606:4700:4700::1001
MTU = 1280
[Peer]
PublicKey = 4
AllowedIPs = 0.0.0.0/0, ::/0
Endpoint = 5
PersistentKeepalive = 25
`

	result, err := generateProfile(&ProfileData{
		PrivateKey: "1",
		Address1:   "2",
		Address2:   "3",
		PublicKey:  "4",
		Endpoint:   "5",
		DNS:        "1.1.1.1, 1.0.0.1, 2606:4700:4700::1111, 2606:4700:4700::1001",
		Keepalive:  25,
	})
	if err != nil {
		t.Fatal(err)
	}

	if expectedResult != result {
		t.Errorf("unexpected profile:\n%s", result)
	}
}

func TestGenerateProfileWithCustomKeepalive(t *testing.T) {
	result, err := generateProfile(&ProfileData{
		PrivateKey: "1",
		Address1:   "2",
		Address2:   "3",
		PublicKey:  "4",
		Endpoint:   "5",
		Keepalive:  60,
	})
	if err != nil {
		t.Fatal(err)
	}

	expected := "PersistentKeepalive = 60\n"
	if !strings.HasSuffix(result, expected) {
		t.Errorf("expected profile to end with %q, got:\n%s", expected, result)
	}
}

func TestGenerateProfileWithCustomDNS(t *testing.T) {
	var expectedResult = `[Interface]
PrivateKey = 1
Address = 2/32, 3/128
DNS = 8.8.8.8, 2001:4860:4860::8888
MTU = 1280
[Peer]
PublicKey = 4
AllowedIPs = 0.0.0.0/0, ::/0
Endpoint = 5
`

	result, err := generateProfile(&ProfileData{
		PrivateKey: "1",
		Address1:   "2",
		Address2:   "3",
		PublicKey:  "4",
		Endpoint:   "5",
		DNS:        "8.8.8.8, 2001:4860:4860::8888",
	})
	if err != nil {
		t.Fatal(err)
	}

	if expectedResult != result {
		t.Errorf("unexpected profile:\n%s", result)
	}
}

func TestGenerateProfileWithoutDNS(t *testing.T) {
	var expectedResult = `[Interface]
PrivateKey = 1
Address = 2/32, 3/128
MTU = 1280
[Peer]
PublicKey = 4
AllowedIPs = 0.0.0.0/0, ::/0
Endpoint = 5
`

	result, err := generateProfile(&ProfileData{
		PrivateKey: "1",
		Address1:   "2",
		Address2:   "3",
		PublicKey:  "4",
		Endpoint:   "5",
		DNS:        "",
	})
	if err != nil {
		t.Fatal(err)
	}

	if expectedResult != result {
		t.Errorf("unexpected profile:\n%s", result)
	}
}
