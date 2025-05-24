package consts

const (
	LabelOtp         = "otp"      // Метка для OTP в keyring
	LabelPassword    = "password" // Метка для пароля в keyring
	LaunchAgentPlist = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" 
  "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>Label</key>
  <string>com.vpn.auto</string>
  <key>ProgramArguments</key>
  <array>
    <string>%s</string>
    <string>daemon</string>
  </array>
  <key>RunAtLoad</key>
  <true/>
  <key>KeepAlive</key>
  <true/>
</dict>
</plist>
` // LaunchAgentPathMacOS - путь к plist-файлу для автозапуска на macOS
)
