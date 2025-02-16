#!/bin/bash
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT="$SCRIPT_DIR/.."
cd "$PROJECT_ROOT" || exit 1

go test -coverprofile=coverage.out ./...
if [ $? -ne 0 ]; then
  echo "Tests failed or no packages to test."
  exit 1
fi

go tool cover -html=coverage.out -o coverage.html

if command -v xdg-open >/dev/null 2>&1; then
  xdg-open coverage.html
elif command -v open >/dev/null 2>&1; then
  open coverage.html
else
  echo "Please open coverage.html manually."
fi
