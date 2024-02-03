.PHONY: setup
setup:
	go install github.com/cosmtrek/air@latest
	docker compose up -d

.PHONY: run
run:
	air