# Pokemon API

A simple REST API for Pokemon data built with Go and Gin.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
  - [Using Swagger UI](#using-swagger-ui)
  - [Regenerating Swagger Documentation](#regenerating-swagger-documentation)
  - [Available Endpoints](#available-endpoints)
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
POKE_API_BASE_URL=https://pokeapi.co/api/v2/pokemon/
```

Available configuration options:

| Variable          | Description                              | Default                              |
|-------------------|------------------------------------------|--------------------------------------|
| POKE_API_BASE_URL | Base URL for the Pokemon API             | https://pokeapi.co/api/v2/pokemon/  |

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

The API is documented using Swagger. This provides interactive documentation where you can explore the API endpoints and even test them directly from your browser.

### Using Swagger UI

After starting the application, you can access the Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

The Swagger UI provides:
- A list of all available endpoints
- Interactive documentation that lets you try out API calls
- Schema information for request and response models
- Detailed descriptions of parameters and response codes

### Regenerating Swagger Documentation

If you make changes to the API endpoints or models, you need to regenerate the Swagger documentation:

```bash
go run github.com/swaggo/swag/cmd/swag init -g cmd/main.go
```

This command will scan your codebase for Swagger annotations and generate updated documentation.

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
│   ├── middleware/      # HTTP middleware examples
│   ├── models/          # Data models
│   ├── repositories/    # Data access layer
│   ├── routes/          # Route definitions
│   └── services/        # Business logic
├── scripts/             # Utility scripts
├── .env                 # Environment variables
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
├── Dockerfile           # Docker configuration
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

For production deployment, consider setting these environment variables:

```
POKE_API_BASE_URL=https://pokeapi.co/api/v2/pokemon/
``` 