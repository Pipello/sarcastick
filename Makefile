.PHONY: fmt
fmt: 
	go mod tidy
	go fmt ./...

.PHONY: refresh
refresh:
	go run cmd/refresh_news/main.go