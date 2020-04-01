package apiserver

import "github.com/Traliaa/http-rest-api/internal/app/store"

//создание структуры конфига
type Config struct {
	BindAddr string `toml:"bind_addr""`
	LogLevel string `tonl:"log_level"`
	Store    *store.Config
}

//новый конфиг
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
