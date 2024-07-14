# Simple Makefile for a Go project
POSTGRESQL_URL=postgres://admin:admin@localhost:5432/neoway?sslmode=disable

# Migration tool
tools:
	@mkdir bin
	@curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-386.tar.gz | tar xvz -C bin

# Build the application
build-app: tools docker-run

build-local:
	@echo "Building..."
	@go build -o bin/main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Create DB container
docker-run:
	@docker compose up -d

# Shutdown DB container
docker-down:
	@docker compose down


# Migrate DB
migrate-up:
	@./bin/migrate -database $(POSTGRESQL_URL) -path db/migrations up

migrate-down:
	@./bin/migrate -database $(POSTGRESQL_URL) -path db/migrations down


# Test the application
test-local:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean-local:
	@echo "Cleaning..."
	@rm -f bin/main
	@rm -f tmp/*
	@rm -f temp/*


# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/air-verse/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: test
test:
	@go test -coverprofile=cover.out ./...

.PHONY: clean
clean:
	@rm -f cover.*