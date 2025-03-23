# Pokemon API

A simple REST API for Pokemon data built with Go and Gin.

## Quick Start

```bash
# Install dependencies
go mod download

# Run the application
go run cmd/main.go

# Access Swagger UI
open http://localhost:8080/swagger/index.html
```

## Features

- REST API with Gin framework
- Random Pokemon endpoint with data from PokeAPI
- In-memory caching
- Swagger documentation

## Configuration

Set in `.env` file:

```
POKE_API_BASE_URL=https://pokeapi.co/api/v2/pokemon/
```

## API Endpoints

| Method | Endpoint               | Description          |
|--------|------------------------|----------------------|
| GET    | /api/v1/pokemon/random | Get a random Pokemon |

## Development

### Regenerate Swagger Docs

```bash
# Using the provided script
./scripts/swagger.sh

# Or manually
go run github.com/swaggo/swag/cmd/swag init -g cmd/main.go
```

### Project Structure

```
├── cmd/                 # Application entry
├── docs/                # Swagger docs
├── internal/            # Private code
│   ├── bootstrap/       # App initialization
│   ├── cache/           # Caching
│   ├── handlers/        # HTTP handlers
│   ├── logger/          # Logging
│   ├── models/          # Data models
│   ├── repositories/    # Data access
│   ├── routes/          # API routes
│   └── services/        # Business logic
└── scripts/             # Utility scripts
```

### Testing

```bash
go test ./...
```

## Docker

```bash
# Build
docker build -t pokemon-api .

# Run
docker run -p 8080:8080 pokemon-api
``` 