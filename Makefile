APP_NAME=dreampicai

run: build
	@./bin/$(APP_NAME)

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss
	@npm install -D daisyui@latest

build:
	npm run build
	@templ generate view
	@go build -o bin/$(APP_NAME) main.go

up: ## Database migration up
	@go run cmd/migrate/main.go up

drop:
	@go run cmd/drop/main.go up

down:
	@go run cmd/migrate/main.go down

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

seed:
	@go run cmd/seed/main.go