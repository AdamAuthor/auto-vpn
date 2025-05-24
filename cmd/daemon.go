package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"auto-vpn/internal/config"
	"auto-vpn/internal/consts"
	"auto-vpn/internal/logs"
	"auto-vpn/internal/models"
	"auto-vpn/internal/secrets"
	"auto-vpn/internal/totp"
	"auto-vpn/internal/utils"
	"auto-vpn/internal/vpn"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Фоновый процесс автоподключения к VPN",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[daemon] Запуск фонового процесса подключения к VPN...")

		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println("Ошибка загрузки конфигурации:", err)
			return
		}

		for {
			if !utils.IsConnected() {
				logs.Logger.Println("[!] VPN не подключён. Пытаемся переподключиться...")

				secretOTP, err := secrets.GetSecret(cfg.Username, consts.LabelOtp)
				if err != nil {
					fmt.Println("Ошибка получения OTP из keyring:", err)
					return
				}
				otpCode, err := totp.GenerateTOTP(secretOTP)
				if err != nil {
					fmt.Println("Ошибка генерации OTP:", err)
					return
				}

				secretPassword, err := secrets.GetSecret(cfg.Username, consts.LabelPassword)

				vpnCfg := models.ConfigVPN{
					Host:     cfg.VPNHost,
					Username: cfg.Username,
					Password: secretPassword,
					TOTPCode: otpCode,
				}

				err = vpn.ConnectVPN(vpnCfg)
				if err != nil {
					fmt.Println("Ошибка подключения к VPN:", err)
				}
			} else {
				logs.Logger.Println("[✓] VPN уже подключён.")
			}
		}

		time.Sleep(30 * time.Second)

	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
