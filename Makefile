include .env
# Simple Makefile for a Go project

# Build the application
all: tools build

tools:
	@curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-386.tar.gz | tar xvz -C bin
	@curl -L https://github.com/a-h/templ/releases/download/v0.2.747/templ_Linux_x86_64.tar.gz | tar xvz -C bin

build:
	@echo "Building..."
	@./bin/templ generate

	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi


# Migrate DB
migrate-up:
	@./bin/migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate-down:
	@./bin/migrate -database ${POSTGRESQL_URL} -path db/migrations down


# Test the application
test:
	@echo "Testing..."
	@echo ${DB_DATABASE}
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main
	@rm -f tmp/*


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

.PHONY: all build run test clean