package config

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"auto-vpn/internal/consts"
)

// EnableAutostart устанавливает автозапуск для текущей ОС
func EnableAutostart(pathToBinary string) error {
	if runtime.GOOS != "darwin" {
		return fmt.Errorf("EnableAutostart поддерживается только для macOS")
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	plistDir := filepath.Join(home, "Library", "LaunchAgents")
	plistPath := filepath.Join(plistDir, "com.vpn.auto.plist")

	err = os.MkdirAll(plistDir, 0755)
	if err != nil {
		return err
	}

	plistContent := fmt.Sprintf(consts.LaunchAgentPlist, pathToBinary)
	err = os.WriteFile(plistPath, []byte(plistContent), 0644)
	if err != nil {
		return err
	}

	// Загружаем plist через launchctl
	cmd := exec.Command("launchctl", "load", plistPath)
	return cmd.Run()
}
