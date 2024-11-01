up:
	docker compose up -d

.PHONY: proto
proto:
	cd proto && buf generate
	cd proto && npx openapi-merge-cli

run: migrate
	go run cmd/rpc/main.go

migrate:
	go run cmd/migrate/main.go

secrets:
	ejson encrypt devops/config.local.json 
