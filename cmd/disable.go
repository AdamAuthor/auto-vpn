package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"auto-vpn/internal/config"
)

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Отключить автоподключение к VPN",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[disable] Отключение автоподключения...")

		if _, err := config.LoadConfig(); err != nil {
			fmt.Println("Ошибка загрузки конфигурации:", err)
			return
		}

		plistPath := os.ExpandEnv("$HOME/Library/LaunchAgents/com.vpn.auto.plist")
		if err := os.Remove(plistPath); err != nil && !os.IsNotExist(err) {
			fmt.Println("Не удалось удалить LaunchAgent:", err)
			return
		}
		fmt.Println("[✓] LaunchAgent удалён:", plistPath)
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
}
