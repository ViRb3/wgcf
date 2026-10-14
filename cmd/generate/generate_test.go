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

func TestKeepaliveFlag(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    int
		wantErr bool
	}{
		{name: "omitted", want: 0},
		{name: "without value", args: []string{"--keepalive"}, wantErr: true},
		{name: "custom (equal syntax)", args: []string{"--keepalive=60"}, want: 60},
		{name: "custom (space-separated syntax)", args: []string{"--keepalive", "60"}, want: 60},
		{name: "short without value", args: []string{"-k"}, wantErr: true},
		{name: "short custom (equal syntax)", args: []string{"-k=60"}, want: 60},
		{name: "short custom (space-separated syntax)", args: []string{"-k", "60"}, want: 60},
		{name: "disabled (equal syntax)", args: []string{"--keepalive=0"}, want: 0},
		{name: "disabled (space-separated syntax)", args: []string{"--keepalive", "0"}, want: 0},
		{name: "invalid value (equal syntax)", args: []string{"--keepalive=invalid"}, wantErr: true},
		{name: "invalid value (space-separated syntax)", args: []string{"--keepalive", "invalid"}, wantErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resetFlag(t, "keepalive")

			err := Cmd.ParseFlags(test.args)
			if (err != nil) != test.wantErr {
				t.Fatalf("ParseFlags(%q) error = %v, wantErr %v", test.args, err, test.wantErr)
			}
			if test.wantErr {
				return
			}
			if keepalive != test.want {
				t.Errorf("keepalive = %d, want %d", keepalive, test.want)
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
