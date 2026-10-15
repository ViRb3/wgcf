package generate

import (
	"slices"
	"testing"

	"github.com/spf13/pflag"
)

func TestValidateKeepalive(t *testing.T) {
	tests := []struct {
		name     string
		interval int
		wantErr  bool
	}{
		{name: "negative", interval: -1, wantErr: true},
		{name: "disabled", interval: 0},
		{name: "default", interval: 25},
		{name: "maximum", interval: maxKeepalive},
		{name: "above maximum", interval: maxKeepalive + 1, wantErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateKeepalive(test.interval)
			if (err != nil) != test.wantErr {
				t.Fatalf("validateKeepalive(%d) error = %v, wantErr %v", test.interval, err, test.wantErr)
			}
		})
	}
}

func TestValidateDns(t *testing.T) {
	tests := []struct {
		name    string
		servers []string
		wantErr bool
	}{
		{name: "nil", wantErr: true},
		{name: "empty", servers: []string{}, wantErr: true},
		{name: "default", servers: defaultDNS},
		{name: "IPv4", servers: []string{"8.8.8.8"}},
		{name: "IPv6", servers: []string{"2001:4860:4860::8888"}},
		{name: "mixed", servers: []string{"8.8.8.8", "2001:4860:4860::8888"}},
		{name: "multiple IPv4", servers: []string{"1.1.1.1", "8.8.8.8"}},
		{name: "invalid IPv4", servers: []string{"8.8.8.999"}, wantErr: true},
		{name: "invalid IPv6", servers: []string{"2001:db8::xyz"}, wantErr: true},
		{name: "hostname", servers: []string{"example.com"}, wantErr: true},
		{name: "IPv6 with zone", servers: []string{"fe80::1%eth0"}, wantErr: true},
		{name: "empty address", servers: []string{""}, wantErr: true},
		{name: "mixed with invalid", servers: []string{"8.8.8.8", "invalid"}, wantErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateDNS(test.servers)
			if (err != nil) != test.wantErr {
				t.Fatalf("validateDns(%q) error = %v, wantErr %v", test.servers, err, test.wantErr)
			}
		})
	}
}

func TestKeepaliveFlag(t *testing.T) {
	flags := Cmd.PersistentFlags()
	flag := flags.Lookup("keepalive")

	defer func() {
		keepalive = 0
		flag.Changed = false
	}()

	keepalive = 0
	if keepalive != 0 {
		t.Error()
	}

	err := flags.Parse([]string{"--keepalive"})
	if err != nil {
		t.Error(err)
	}
	if keepalive != 25 {
		t.Error()
	}

	keepalive = 0
	flag.Changed = false
	err = flags.Parse([]string{"--keepalive=60"})
	if err != nil {
		t.Error(err)
	}
	if keepalive != 60 {
		t.Error()
	}
}

func TestDNSFlags(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantDNS      []string
		wantNoDNS    bool
		wantParseErr bool
		wantGroupErr bool
	}{
		{name: "omitted", wantDNS: defaultDNS},
		{name: "custom IPv4", args: []string{"--dns=8.8.8.8"}, wantDNS: []string{"8.8.8.8"}},
		{name: "custom IPv6", args: []string{"--dns=2001:4860:4860::8888"}, wantDNS: []string{"2001:4860:4860::8888"}},
		{name: "mixed", args: []string{"--dns=8.8.8.8,2001:4860:4860::8888"}, wantDNS: []string{"8.8.8.8", "2001:4860:4860::8888"}},
		{name: "comma separated", args: []string{"--dns=8.8.8.8,1.1.1.1"}, wantDNS: []string{"8.8.8.8", "1.1.1.1"}},
		{name: "repeated", args: []string{"--dns=8.8.8.8", "--dns=1.1.1.1"}, wantDNS: []string{"8.8.8.8", "1.1.1.1"}},
		{name: "short", args: []string{"-d", "8.8.8.8"}, wantDNS: []string{"8.8.8.8"}},
		{name: "no DNS", args: []string{"--no-dns"}, wantDNS: defaultDNS, wantNoDNS: true},
		{name: "no DNS explicitly false", args: []string{"--no-dns=false"}, wantDNS: defaultDNS},
		{name: "empty DNS", args: []string{"--dns="}},
		{name: "missing DNS value", args: []string{"--dns"}, wantParseErr: true},
		{name: "conflicting flags", args: []string{"--dns=8.8.8.8", "--no-dns"}, wantDNS: []string{"8.8.8.8"}, wantNoDNS: true, wantGroupErr: true},
		{name: "conflicting flags reversed", args: []string{"--no-dns", "--dns=8.8.8.8"}, wantDNS: []string{"8.8.8.8"}, wantNoDNS: true, wantGroupErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resetDNSFlags(t)

			err := Cmd.ParseFlags(test.args)
			if (err != nil) != test.wantParseErr {
				t.Fatalf("ParseFlags(%q) error = %v, wantParseErr %v", test.args, err, test.wantParseErr)
			}
			if test.wantParseErr {
				return
			}

			err = Cmd.ValidateFlagGroups()
			if (err != nil) != test.wantGroupErr {
				t.Fatalf("ValidateFlagGroups() error = %v, wantGroupErr %v", err, test.wantGroupErr)
			}
			if !slices.Equal(dnsServers, test.wantDNS) {
				t.Errorf("dnsServers = %q, want %q", dnsServers, test.wantDNS)
			}
			if noDNS != test.wantNoDNS {
				t.Errorf("noDNS = %t, want %t", noDNS, test.wantNoDNS)
			}
		})
	}
}

func resetDNSFlags(t *testing.T) {
	t.Helper()

	dnsFlag := Cmd.PersistentFlags().Lookup("dns")
	noDNSFlag := Cmd.PersistentFlags().Lookup("no-dns")

	oldValue := dnsFlag.Value
	oldDNS := slices.Clone(dnsServers)
	oldNoDNS := noDNS
	oldDNSChanged := dnsFlag.Changed
	oldNoDNSChanged := noDNSFlag.Changed

	t.Cleanup(func() {
		dnsFlag.Value = oldValue
		dnsFlag.Changed = oldDNSChanged
		noDNSFlag.Changed = oldNoDNSChanged
		dnsServers = oldDNS
		noDNS = oldNoDNS
	})

	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flags.StringSliceVar(&dnsServers, "dns", defaultDNS, "")

	dnsFlag.Value = flags.Lookup("dns").Value
	dnsFlag.Changed = false
	noDNS = false
	noDNSFlag.Changed = false
}
