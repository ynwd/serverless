build:
	go run cmd/static/main.go
	git add .
	git commit -m "new date commit"
	git push origin main