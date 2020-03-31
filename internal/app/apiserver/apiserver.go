package apiserver

//создание структуры сервера
type APIServer struct {
	config *Config
}

//инициализация структура
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

//функция старта сервера
func (s *APIServer) Start() error {
	return nil
}
