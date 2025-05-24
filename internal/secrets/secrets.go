package secrets

import (
	"fmt"

	"github.com/zalando/go-keyring"
)

const service = "vpn-auto"

// SaveSecret сохраняет значение (пароль или OTP) в keyring под заданной меткой
func SaveSecret(username, label, value string) error {
	key := fmt.Sprintf("%s:%s", username, label)
	return keyring.Set(service, key, value)
}

// GetSecret извлекает значение (пароль или OTP) из keyring
func GetSecret(username, label string) (string, error) {
	key := fmt.Sprintf("%s:%s", username, label)
	return keyring.Get(service, key)
}

// DeleteSecret удаляет значение из keyring
func DeleteSecret(username, label string) error {
	key := fmt.Sprintf("%s:%s", username, label)
	return keyring.Delete(service, key)
}
