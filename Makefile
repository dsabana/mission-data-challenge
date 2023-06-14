PG_HOST?=localhost
PG_PORT?=5432
PG_USER?=mission_data
PG_PASSWORD?=secret
PG_DB?=mission_data
PG_SCHEMA?=journal

.PHONY: pg
pg:
	docker ps -a -q --filter "name=mission-data-db" | grep -q . && docker rm -f mission-data-db || echo Awesome!! No DB present.
	docker run -d --name mission-data-db -p 5432:5432 -e POSTGRES_USER=$(PG_USER) -e POSTGRES_PASSWORD=$(PG_PASSWORD) postgres:14.1-alpine
	sleep 5
	docker exec -it mission-data-db psql -U$(PG_USER) -a $(PG_USER) -c 'CREATE SCHEMA IF NOT EXISTS $(PG_SCHEMA);'

.PHONY: migrate-up
migrate-up:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	migrate -path db/migrations -database "postgresql://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DB)?sslmode=disable&search_path=$(PG_SCHEMA)" -verbose up

.PHONY: local-db
local-db: pg migrate-up

.PHONY: run
run:
	go run cmd/main.go

