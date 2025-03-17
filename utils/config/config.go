package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	IsInit      bool         `json:"isInit"`
	CacheEngine CacheEngine  `json:"cacheEngine"`
	Redis       RedisConfig  `json:"redis"`
	Db          DbConfig     `json:"db"`
	Secret      SecretConfig `json:"secret"`
}

func (c *Config) InitConfig() error {
	file, err := os.Open("./config.json")
	if err != nil {
		println("err", err)
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) SaveConfig() error {
	file, err := os.Create("./config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		return err
	}
	return nil
}
