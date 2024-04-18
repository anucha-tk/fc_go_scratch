include .env

init:
	@echo "== 👩‍🌾 init =="
	brew install pre-commit
	brew install golangci-lint
	brew upgrade golangci-lint

	@echo "== pre-commit setup =="
	pre-commit install

precommit.rehooks:
	pre-commit autoupdate
	pre-commit install --install-hooks
	pre-commit install --hook-type commit-msg

ci.lint:
	@echo "== 🙆 ci.linter =="
	golangci-lint run -v ./... --fix

run-dev:
	docker compose up -d
stop:
	docker compose stop
down:
	docker compose down
exec:
	docker exec -it go_scratch_postgres bin/bash
goose-up:
	cd sql/schema && goose postgres "postgresql://$(DATABASE_USERNAME):$(DATABASE_PASSWORD)@localhost:5432/$(DATABASE_NAME)?sslmode=disable" up
goose-down:
	cd sql/schema && goose postgres "postgresql://$(DATABASE_USERNAME):$(DATABASE_PASSWORD)@localhost:5432/$(DATABASE_NAME)?sslmode=disable" down
sqlc-gen:
	sqlc generate
