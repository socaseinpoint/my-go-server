# Установка зависимостей
install:
	go mod tidy

# Запуск приложения
run:
	go run main.go

# Генерация Swagger документации
swagger:
	swag init --dir ./,./routes --parseDependency

# Тестирование
test:
	go test ./...

# Сборка
build:
	go build -o my-go-server main.go

# Запуск приложения с перезагрузкой
run-watch:
	air
