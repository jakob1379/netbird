package main

import (
	"fmt"
	"net/netip"
	"os"
	"path/filepath"
	"strings"

	sshconfig "github.com/netbirdio/netbird/client/ssh/config"
)

func mark(name string) {
	dir := os.Getenv("NB_E2E_MARKERS")
	if dir == "" {
		return
	}
	_ = os.WriteFile(filepath.Join(dir, name), []byte(strings.Join(os.Args, " ")), 0o644)
}

func main() {
	if len(os.Args) >= 3 && os.Args[1] == "ssh" {
		switch os.Args[2] {
		case "detect":
			mark("detect")
			os.Exit(0)
		case "proxy":
			mark("proxy")
			os.Exit(1)
		}
	}

	peers := []sshconfig.PeerSSHInfo{{
		Hostname: "peer1",
		IP:       netip.MustParseAddr("100.64.0.1"),
		FQDN:     "peer1.nb.internal",
	}}
	if err := sshconfig.New().SetupSSHClientConfig(peers); err != nil {
		fmt.Fprintln(os.Stderr, "generate:", err)
		os.Exit(1)
	}
}
