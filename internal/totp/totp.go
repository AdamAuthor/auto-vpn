package totp

import (
	"time"

	totplib "github.com/pquerna/otp/totp"
)

func GenerateTOTP(secret string) (string, error) {
	return totplib.GenerateCode(secret, time.Now())
}
