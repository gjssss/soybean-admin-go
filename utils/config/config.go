package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	CacheEngine CacheEngine  `json:"cache_engine"`
	Redis       RedisConfig  `json:"redis"`
	Db          DbConfig     `json:"db"`
	Secret      SecretConfig `json:"secret"`
}

func InitConfig() (Config, error) {
	file, err := os.Open("./config.json")
	if err != nil {
		println("err", err)
		return Config{}, err
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
