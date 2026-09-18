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
