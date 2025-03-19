package config

type RedisConfig struct {
	Host     string  `json:"host"`
	DB       int     `json:"db"`
	Password *string `json:"password"`
	Username *string `json:"username"`
}
