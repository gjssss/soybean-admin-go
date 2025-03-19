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
	S3          S3Config     `json:"s3"`
}

// S3Config is a struct that stores the S3 configuration
type S3Config struct {
	Endpoint     string `json:"endpoint"`
	Bucket       string `json:"bucket"`
	PublicDomain string `json:"publicDomain"`
	AccessKey    string `json:"accessKey,omitempty"`
	SecretKey    string `json:"secretKey,omitempty"`
	Region       string `json:"region,omitempty"`
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
