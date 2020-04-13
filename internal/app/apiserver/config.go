package apiserver

//создание структуры конфига
type Config struct {
	BindAddr    string `toml:"bind_addr""`
	LogLevel    string `toml:"log_level"`
	DatabaseURl string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

//новый конфиг
func NewConfig() *Config {
	return &Config{
		BindAddr: ":80",
		LogLevel: "debug",
	}
}
