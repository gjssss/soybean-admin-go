package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

type Secret struct {
	JwtKey    string
	StartTime int64
}

type Config struct {
	DB     *DBConfig
	Secret *Secret
}

func (r *Config) Init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	r.DB = &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}
	r.Secret = &Secret{
		JwtKey:    os.Getenv("JWT_KEY"),
		StartTime: time.Now().Unix(),
	}
}
