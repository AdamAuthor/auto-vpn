package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"auto-vpn/internal/models"
)

func configPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".vpn_auto_config.json"), nil
}

func SaveConfig(cfg models.Config) error {
	path, err := configPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0600)
}

func LoadConfig() (models.Config, error) {
	path, err := configPath()
	if err != nil {
		return models.Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return models.Config{}, errors.New("конфигурационный файл не найден, выполните vpn-auto setup")
		}
		return models.Config{}, err
	}

	var cfg models.Config
	err = json.Unmarshal(data, &cfg)
	return cfg, err
}
