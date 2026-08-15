package generate

import "testing"

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
