build:
	go run cmd/static/main.go
	go build -o ./tmp/main cmd/main.go