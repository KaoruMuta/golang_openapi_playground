.PHONY: setup
setup:
	go install github.com/cosmtrek/air@latest

.PHONY: run
run: setup
	air

.PHONY: up
up: setup
	docker compose up -d

.PHONY: reset
reset:
	docker compose down -v
	docker rmi golang_openapi_playground-app