.PHONY: setup
setup:
	go install github.com/cosmtrek/air@latest
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
	go mod tidy

.PHONY: generate-oapi-codegen
generate-oapi-codegen:
	mkdir -p oapi_codegen/generated/
	oapi-codegen -generate types -package generated openapi.yml > oapi_codegen/generated/model_gen.go
	oapi-codegen -generate server -package generated openapi.yml > oapi_codegen/generated/api_server_gen.go
	oapi-codegen -generate client -package generated openapi.yml > oapi_codegen/generated/api_client_gen.go

.PHONY: generate-openapi-generator
generate-openapi-generator:
	mkdir -p openapi_generator/generated/
	docker run --rm \
		-v ${PWD}:/app openapitools/openapi-generator-cli generate \
		-i /app/openapi.yml \
		-g go \
		-o /app/openapi_generator/generated

.PHONY: generate
generate: generate-oapi-codegen generate-openapi-generator

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