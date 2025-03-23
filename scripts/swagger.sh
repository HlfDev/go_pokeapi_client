#!/bin/bash

# Regenerate Swagger documentation
echo "Regenerating Swagger documentation..."
go run github.com/swaggo/swag/cmd/swag init -g cmd/main.go

if [ $? -eq 0 ]; then
    echo "✅ Swagger documentation regenerated successfully!"
    echo "Access Swagger UI at: http://localhost:8080/swagger/index.html"
else
    echo "❌ Failed to regenerate Swagger documentation"
    exit 1
fi 