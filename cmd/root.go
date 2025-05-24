package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vpn-auto",
	Short: "Автоматическое подключение к Cisco AnyConnect VPN",
	Long:  `vpn-auto — CLI-утилита для автоподключения к VPN через Cisco AnyConnect`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Используйте одну из доступных команд: setup, daemon, disable, status")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
