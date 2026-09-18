package util

import "testing"

func TestIsWarpLicenseKey(t *testing.T) {
	tests := []struct {
		value string
		want  bool
	}{
		{"ABCDEFGH-12345678-abcdefgh", true},
		{"ABCDEFGH-12345678", false},
		{"ABCDEFGH-12345678-abcdefg-", false},
		{"ABCDEFGH_12345678_abcdefgh", false},
	}
	for _, test := range tests {
		if got := IsWarpLicenseKey(test.value); got != test.want {
			t.Errorf("IsWarpLicenseKey(%q) = %t, want %t", test.value, got, test.want)
		}
	}
}
