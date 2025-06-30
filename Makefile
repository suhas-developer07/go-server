include .env

MIGRATION_PATH = ./migrations

.phony: migrate-create

migrate-create:
     @migrate-create -seq -ext sql -dir ${MIGRATION_PATH} $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
    @migrate -path=$(MIGRATION_PATH) -database=$(POSTGRES_URL) up 

migrate-down:
    @migrate-path=$(MIGRATION_PATH) -database=$(POSTGRES_URL) down $(filter-out $@,$(MAKECMDGOALS))