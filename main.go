package main

import (
	"auto-vpn/cmd"
	"auto-vpn/internal/logs"
)

func main() {
	logs.Init()
	defer logs.Close()
	cmd.Execute()
}
