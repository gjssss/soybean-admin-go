package config

type SecretConfig struct {
	JwtKey    string `json:"jwt"`
	StartTime int64  `json:"-"`
}
