package models

type Config struct {
	OS        string `json:"os"`
	VPNHost   string `json:"vpn_host"`
	Username  string `json:"username"`
	AutoStart bool   `json:"auto_start"`
}
