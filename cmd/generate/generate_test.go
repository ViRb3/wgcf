package generate

import "testing"

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

func TestValidateMTU(t *testing.T) {
	tests := []struct {
		name    string
		mtu     int
		wantErr bool
	}{
		{name: "negative", mtu: -1, wantErr: true},
		{name: "zero", mtu: 0, wantErr: true},
		{name: "minimum accepted", mtu: 1},
		{name: "default", mtu: 1280},
		{name: "custom", mtu: 1500},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateMTU(test.mtu)
			if (err != nil) != test.wantErr {
				t.Fatalf("validateMTU(%d) error = %v, wantErr %v", test.mtu, err, test.wantErr)
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

func TestMTUFlags(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantMTU      int
		wantNoMTU    bool
		wantParseErr bool
		wantGroupErr bool
	}{
		{name: "omitted", wantMTU: 1280},
		{name: "custom", args: []string{"--mtu=1500"}, wantMTU: 1500},
		{name: "short", args: []string{"-m", "1400"}, wantMTU: 1400},
		{name: "no MTU", args: []string{"--no-mtu"}, wantMTU: 1280, wantNoMTU: true},
		{name: "no MTU explicitly false", args: []string{"--no-mtu=false"}, wantMTU: 1280},
		{name: "missing MTU value", args: []string{"--mtu"}, wantParseErr: true},
		{name: "invalid MTU value", args: []string{"--mtu=invalid"}, wantParseErr: true},
		{name: "conflicting flags", args: []string{"--mtu=1500", "--no-mtu"}, wantMTU: 1500, wantNoMTU: true, wantGroupErr: true},
		{name: "conflicting flags reversed", args: []string{"--no-mtu", "--mtu=1500"}, wantMTU: 1500, wantNoMTU: true, wantGroupErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resetFlag(t, "mtu")
			resetFlag(t, "no-mtu")

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
			if mtu != test.wantMTU {
				t.Errorf("mtu = %d, want %d", mtu, test.wantMTU)
			}
			if noMTU != test.wantNoMTU {
				t.Errorf("noMTU = %t, want %t", noMTU, test.wantNoMTU)
			}
		})
	}
}

func resetFlag(t *testing.T, name string) {
	t.Helper()
	flag := Cmd.PersistentFlags().Lookup(name)
	if flag == nil {
		t.Fatalf("flag %q not found", name)
	}
	value, changed := flag.Value.String(), flag.Changed
	t.Cleanup(func() {
		if err := flag.Value.Set(value); err != nil {
			t.Errorf("restore flag %q: %v", name, err)
		}
		flag.Changed = changed
	})
	if err := flag.Value.Set(flag.DefValue); err != nil {
		t.Fatalf("reset flag %q: %v", name, err)
	}
	flag.Changed = false
}
