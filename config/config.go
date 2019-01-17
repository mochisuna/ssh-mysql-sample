package config

import "github.com/BurntSushi/toml"

func New(config *Config, configPath string) error {
	_, err := toml.DecodeFile(configPath, config)
	return err
}

type Config struct {
	SSH SSH `toml:"ssh"`
	DB  DB  `toml:"db"`
}

type DB struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
}

type SSH struct {
	Key  string `toml:"key"`
	Host string `toml:"host"`
	Port string `toml:"port"`
	User string `toml:"user"`
}
