package vpn

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"

	"auto-vpn/internal/models"
)

func ConnectVPN(cfg models.ConfigVPN) error {
	var vpnCmd string
	switch runtime.GOOS {
	case "darwin", "linux":
		vpnCmd = "/opt/cisco/anyconnect/bin/vpn"
	case "windows":
		vpnCmd = "vpncli.exe"
	default:
		return fmt.Errorf("неподдерживаемая ОС: %s", runtime.GOOS)
	}

	cmd := exec.Command(vpnCmd, "connect", cfg.Host)
	stdin := &bytes.Buffer{}
	stdin.WriteString(fmt.Sprintf("%s\n", cfg.Username))
	stdin.WriteString(fmt.Sprintf("%s\n", cfg.Password+cfg.TOTPCode))
	cmd.Stdin = stdin

	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	return err
}
