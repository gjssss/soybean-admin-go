package config

type DbConfig struct {
	Type     string `json:"type"` // 数据库类型: mysql, postgres
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	SslMode  string `json:"sslmode"`
	Timezone string `json:"timezone"`
}
