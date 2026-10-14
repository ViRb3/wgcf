package cmd

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

func TestVersionDevOutputFormat(t *testing.T) {
	var output bytes.Buffer

	RootCmd.SetOut(&output)
	RootCmd.SetArgs([]string{"--version"})

	t.Cleanup(func() {
		RootCmd.SetOut(nil)
		RootCmd.SetArgs(nil)
	})

	if err := RootCmd.Execute(); err != nil {
		t.Fatal(err)
	}

	want := version + "\n"
	if got := output.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestVersionSetOnBuild(t *testing.T) {
	const want = "v5.5.5\n"

	binary := filepath.Join(t.TempDir(), "wgcf")
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}

	build := exec.Command(
		"go", "build",
		"-ldflags", "-X github.com/ViRb3/wgcf/v2/cmd.version=v5.5.5",
		"-o", binary,
		"..",
	)

	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("go build failed: %v\n%s", err, string(output))
	}

	output, err := exec.Command(binary, "--version").Output()
	if err != nil {
		t.Fatalf("go version failed: %v\n%s", err, string(output))
	}

	if string(output) != want {
		t.Errorf("got %q, want %q", string(output), want)
	}
}
