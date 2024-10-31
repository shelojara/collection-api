up:
	docker compose up -d

.PHONY: proto
proto:
	cd proto && buf generate

run: migrate
	go run cmd/rpc/main.go

migrate:
	go run cmd/migrate/main.go
