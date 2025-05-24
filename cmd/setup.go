package cmd

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"

	"auto-vpn/internal/config"
	"auto-vpn/internal/models"
	"auto-vpn/internal/secrets"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Настроить VPN и автоподключение",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[setup] Запуск мастера настройки...")
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Определяем вашу операционную систему...")
		osName := runtime.GOOS
		fmt.Printf("Обнаружена ОС: %s\n", osName)

		fmt.Print("Введите адрес VPN-сервера: ")
		host, _ := reader.ReadString('\n')

		fmt.Print("Введите имя пользователя: ")
		username, _ := reader.ReadString('\n')

		fmt.Print("Введите пароль: ")
		password, _ := reader.ReadString('\n')

		fmt.Print("Введите секретный ключ для Google Authenticator: ")
		secret, _ := reader.ReadString('\n')

		fmt.Print("Разрешить автозапуск при старте системы? (yes/no): ")
		auto, _ := reader.ReadString('\n')

		cfg := models.Config{
			OS:        osName,
			VPNHost:   strings.TrimSpace(host),
			Username:  strings.TrimSpace(username),
			AutoStart: strings.ToLower(strings.TrimSpace(auto)) == "yes",
		}

		err := secrets.SaveSecret(cfg.Username, "password", password)
		if err != nil {
			fmt.Println("Ошибка сохранения пароля в keyring:", err)
			return
		}

		err = secrets.SaveSecret(cfg.Username, "otp", secret)
		if err != nil {
			fmt.Println("Ошибка сохранения OTP в keyring:", err)
			return
		}

		err = config.SaveConfig(cfg)
		if err != nil {
			fmt.Println("Ошибка сохранения конфигурации:", err)
			return
		}

		fmt.Println("Настройка завершена. Конфигурация сохранена.")

		binPath, _ := os.Executable()
		binPath = strings.TrimSpace(binPath)

		if cfg.AutoStart {
			err := config.EnableAutostart(binPath)
			if err != nil {
				fmt.Println("Не удалось включить автозапуск:", err)
			} else {
				fmt.Println("[✓] Автозапуск активирован через LaunchAgent")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
