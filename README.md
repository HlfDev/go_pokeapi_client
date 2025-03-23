# Pokemon API

A simple REST API for Pokemon data built with Go and Gin.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [Development](#development)
- [Testing](#testing)
- [Deployment](#deployment)

## Features

- REST API built with Gin framework
- Random Pokemon endpoint that fetches data from PokeAPI
- In-memory caching
- Structured logging
- API documentation with Swagger

## Requirements

- Go 1.23 or higher
- Git

## Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/pokeapi.git
cd pokeapi
```

Install dependencies:

```bash
go mod download
```

## Configuration

The application uses environment variables for configuration. You can set these in a `.env` file in the root directory.

Example `.env` file:

```
PORT=8080
GIN_MODE=debug
SWAGGER_HOST=localhost:8080
```

Available configuration options:

| Variable     | Description                             | Default       |
|-------------|-----------------------------------------|---------------|
| PORT        | Port the server will listen on          | 8080          |
| GIN_MODE    | Gin mode (debug or release)             | debug         |
| SWAGGER_HOST| Host used in Swagger documentation      | localhost:8080|
| TEST_MODE   | Set to 'true' for testing               | false         |

## Running the Application

### Development Mode

Run the application in development mode:

```bash
go run cmd/main.go
```

### Build and Run

Build the application:

```bash
go build -o pokemon-api cmd/main.go
```

Run the compiled binary:

```bash
./pokemon-api
```

## API Documentation

The API is documented using Swagger. After starting the application, you can access the Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

### Regenerating Swagger Documentation

If you make changes to the API endpoints or models, you need to regenerate the Swagger documentation:

```bash
go run github.com/swaggo/swag/cmd/swag init -g cmd/main.go
```

### Available Endpoints

| Method | Endpoint                 | Description            |
|--------|--------------------------|------------------------|
| GET    | /api/v1/pokemon/random   | Get a random Pokemon   |

## Development

### Project Structure

```
.
├── cmd/                 # Application entry points
│   └── main.go          # Main application entry
├── docs/                # Generated Swagger documentation
├── internal/            # Private application code
│   ├── bootstrap/       # Application initialization
│   ├── cache/           # Caching implementation
│   ├── handlers/        # HTTP handlers
│   ├── logger/          # Logging configuration
│   ├── middleware/      # HTTP middleware
│   ├── models/          # Data models
│   ├── repositories/    # Data access layer
│   ├── routes/          # Route definitions
│   └── services/        # Business logic
├── scripts/             # Utility scripts
├── .env                 # Environment variables
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
└── README.md            # This file
```

## Testing

Run all tests:

```bash
go test ./...
```

## Deployment

### Docker

Build the Docker image:

```bash
docker build -t pokemon-api .
```

Run the container:

```bash
docker run -p 8080:8080 pokemon-api
```

### Environment Variables for Production

For production deployment, set the following environment variables:

```
GIN_MODE=release
PORT=8080
SWAGGER_HOST=your-production-domain.com
``` 