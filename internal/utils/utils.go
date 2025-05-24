package utils

import (
	"os/exec"
)

func IsConnected() bool {
	cmd := exec.Command("ping", "-c", "1", "-W", "1", "8.8.8.8")
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
