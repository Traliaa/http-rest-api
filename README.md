1. Подключени баз
  изменить файл apiserver.toml и прописать правильный конфиг
2. Создать миграцию баз.
    migrate -patg migrations -database "postgres://loaclhost/restapi_dev?sslmode=disable" up
    
    
