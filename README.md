# Gin Upload File
Upload large files with Golang

## How to run Golang App
Build App and run migration-up
```bash
make build-app
make migrate-up
```

Open the browser at
```bash
http://localhost:8080/
```

## Tables
```bash
CREATE TABLE IF NOT EXISTS customers(
    id SERIAL PRIMARY KEY,
    cpf VARCHAR (14) NOT NULL,
    private INT NOT NULL,
    incomplete INT NOT NULL,
    date_last_purchase TIMESTAMP,
    average_ticket DECIMAL(10,2) DEFAULT(0.00),
    last_purchase_ticket DECIMAL(10,2) DEFAULT(0.00),
    most_frequent_store VARCHAR(18),
    last_purchase_store VARCHAR(18),
    created_at TIMESTAMP DEFAULT 'now()'
);
```

## MakeFile
### Migration tool
```bash
make tools
```

### Build the application
```bash
make build-app
make build-local
```

### Run the application
```bash
make run
```

### Create DB container
```bash
make docker-run
```

### Shutdown DB container
```bash
make docker-down
```

### Migrate DB
```bash
make migrate-up
make migrate-down
```

### Test the application
```bash
make test-local
```

### Clean the binary
```bash
make clean-local
```

### Live Reload
```bash
make watch
```

### Continuous Integration
```bash
.PHONY: test
test:
	@go test -coverprofile=cover.out ./...

.PHONY: clean
clean:
	@rm -f cover.*
```