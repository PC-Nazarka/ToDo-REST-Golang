include .env
export

.PHONY: help

compose-up: ### Run docker-compose
	docker-compose up --build
.PHONY: compose-up

compose-down: ### Down docker-compose
	docker-compose down --remove-orphans
.PHONY: compose-down

swag-v1: ### swag init
	swag init --dir internal/handler --generalInfo handler.go  --output docs --parseDependency --parseInternal
.PHONY: swag-v1

run: swag-v1 ### swag run
	go mod tidy && go mod download && \
	DISABLE_SWAGGER_HTTP_HANDLER='' GIN_MODE=debug CGO_ENABLED=0 go run -tags migrate ./cmd/app
.PHONY: run

docker-rm-volume: ### remove docker volume
	docker volume down -v
.PHONY: docker-rm-volume

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'migrate_name'
.PHONY: migrate-create

migrate-up: ### migration up
	migrate -path ./migrations -database '$(POSTGRES_URL)' up
.PHONY: migrate-up

migrate-down: ### migration up
	migrate -path ./migrations -database '$(POSTGRES_URL)' down
.PHONY: migrate-down