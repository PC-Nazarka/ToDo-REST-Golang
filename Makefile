include .env
export

compose-up:
	docker-compose up --build
.PHONY: compose-up

compose-down:
	docker-compose down --remove-orphans
.PHONY: compose-down

swag-v1:
	swag init --dir internal/handler --generalInfo handler.go  --output docs --parseDependency --parseInternal
.PHONY: swag-v1

docker-rm-volume:
	docker volume down -v
.PHONY: docker-rm-volume

migrate-create:
	migrate create -ext sql -dir migrations 'migrate_name'
.PHONY: migrate-create

migrate-up:
	migrate -path ./migrations -database '$(POSTGRES_URL)' up
.PHONY: migrate-up

migrate-down:
	migrate -path ./migrations -database '$(POSTGRES_URL)' down
.PHONY: migrate-down
