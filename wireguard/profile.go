package wireguard

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

var profileTemplate = `[Interface]
PrivateKey = {{ .PrivateKey }}
Address = {{ .Address1 }}/32
Address = {{ .Address2 }}/128
DNS = 1.1.1.1
MTU = 1280
[Peer]
PublicKey = {{ .PublicKey }}
AllowedIPs = 0.0.0.0/5, 8.0.0.0/7, 11.0.0.0/8, 12.0.0.0/6, 16.0.0.0/4, 32.0.0.0/3, 64.0.0.0/2, 128.0.0.0/3, 160.0.0.0/5, 168.0.0.0/6, 172.0.0.0/12, 172.32.0.0/11, 172.64.0.0/10, 172.128.0.0/9, 173.0.0.0/8, 174.0.0.0/7, 176.0.0.0/4, 192.0.0.0/9, 192.128.0.0/11, 192.160.0.0/13, 192.169.0.0/16, 192.170.0.0/15, 192.172.0.0/14, 192.176.0.0/12, 192.192.0.0/10, 193.0.0.0/8, 194.0.0.0/7, 196.0.0.0/6, 200.0.0.0/5, 208.0.0.0/4, ::/0, 193.138.218.74/32
Endpoint = {{ .Endpoint }}`

type Profile struct {
	profileString string
}

type ProfileData struct {
	PrivateKey string
	Address1   string
	Address2   string
	PublicKey  string
	Endpoint   string
}

func NewProfile(data *ProfileData) (*Profile, error) {
	profileString, err := generateProfile(data)
	if err != nil {
		return nil, err
	}
	return &Profile{profileString: profileString}, nil
}

func generateProfile(data *ProfileData) (string, error) {
	t, err := template.New("").Parse(profileTemplate)
	if err != nil {
		return "", err
	}
	var result bytes.Buffer
	if err := t.Execute(&result, data); err != nil {
		return "", err
	}
	return result.String(), nil
}

func (p *Profile) Save(profileFile string) error {
	return ioutil.WriteFile(profileFile, []byte(p.profileString), 0600)
}
