package apiserver

//создание структуры конфига
type Config struct {
	BindAddr string `toml:"bind_addr""`
}

//новый конфиг
func NewConfig() *Config {
	return &Config{
		BindAddr: "*:8080",
	}
}
