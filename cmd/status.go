package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"auto-vpn/internal/utils"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Показать текущий статус VPN",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[status] Проверка текущего VPN-соединения...")

		connected := utils.IsConnected()

		if connected {
			fmt.Println("[✓] VPN подключён")
		} else {
			fmt.Println("[✗] VPN не подключён")
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
