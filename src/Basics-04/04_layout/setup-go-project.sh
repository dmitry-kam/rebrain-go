#!/bin/bash

# Go Project Structure Setup Script
# Creates a standard Go project directory structure

set -e

# Check if project name is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <project-name>"
    echo "Example: $0 my-awesome-app"
    exit 1
fi

PROJECT_NAME=$1

# Create project root directory
mkdir -p "$PROJECT_NAME"
cd "$PROJECT_NAME"

echo "Creating Go project structure for: $PROJECT_NAME"

# Core directories
mkdir -p cmd/"$PROJECT_NAME"
mkdir -p internal/{handlers,routes,models,middleware,database}
mkdir -p pkg/{utils,logger,httputil}

# API & Web
mkdir -p api/{swagger,proto}
mkdir -p web/{static/{css,js,images},templates}

# Configuration & Scripts
mkdir -p config
mkdir -p scripts

# Build & Deployment
mkdir -p build/{ci,docker}
mkdir -p deployments/{docker-compose,kubernetes}

# Testing & Documentation
mkdir -p test/{integration,unit}
mkdir -p docs

# Utilities & Assets
mkdir -p tools
mkdir -p examples
mkdir -p third_party
mkdir -p assets/{images,icons}

# Create initial files
cat > cmd/"$PROJECT_NAME"/main.go << 'EOF'
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello from " + "PROJECT_NAME_PLACEHOLDER" + "!")
	log.Println("Application started")
}
EOF

# Replace placeholder with actual project name
sed -i.bak "s/PROJECT_NAME_PLACEHOLDER/$PROJECT_NAME/g" cmd/"$PROJECT_NAME"/main.go
rm cmd/"$PROJECT_NAME"/main.go.bak 2>/dev/null || true

# Create go.mod
cat > go.mod << EOF
module github.com/yourusername/$PROJECT_NAME

go 1.21
EOF

# Create .gitignore
cat > .gitignore << 'EOF'
# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib
bin/
dist/

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool
*.out

# Go workspace file
go.work

# Vendor directory
vendor/

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db

# Environment variables
.env
.env.local
EOF

# Create README
cat > README.md << EOF
# $PROJECT_NAME

## Project Structure

\`\`\`
$PROJECT_NAME/
├── cmd/                    # Application entry points
├── internal/              # Private application code
├── pkg/                   # Public library code
├── api/                   # API definitions (Swagger, Proto)
├── web/                   # Web assets and templates
├── config/                # Configuration files
├── scripts/               # Build and utility scripts
├── build/                 # Build configurations
├── deployments/           # Deployment configurations
├── test/                  # Additional test files
├── docs/                  # Documentation
├── tools/                 # Support tools
├── examples/              # Usage examples
├── third_party/           # External tools and utilities
└── assets/                # Images, icons, logos
\`\`\`

## Getting Started

\`\`\`bash
# Build the application
go build -o bin/$PROJECT_NAME ./cmd/$PROJECT_NAME

# Run the application
./bin/$PROJECT_NAME

# Run tests
go test ./...
\`\`\`

## Development

TODO: Add development instructions

## License

TODO: Add license information
EOF

# Create a sample Dockerfile
cat > build/docker/Dockerfile << EOF
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/$PROJECT_NAME ./cmd/$PROJECT_NAME

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/bin/$PROJECT_NAME .

CMD ["./$PROJECT_NAME"]
EOF

# Create a sample docker-compose.yml
cat > deployments/docker-compose/docker-compose.yml << EOF
version: '3.8'

services:
  app:
    build:
      context: ../..
      dockerfile: build/docker/Dockerfile
    ports:
      - "8080:8080"
    environment:
      - ENV=development
    restart: unless-stopped
EOF

# Create a sample Makefile
cat > Makefile << EOF
.PHONY: build run test clean docker-build docker-run

APP_NAME=$PROJECT_NAME
BUILD_DIR=bin

build:
	@echo "Building..."
	@go build -o \$(BUILD_DIR)/\$(APP_NAME) ./cmd/\$(APP_NAME)

run: build
	@echo "Running..."
	@./\$(BUILD_DIR)/\$(APP_NAME)

test:
	@echo "Testing..."
	@go test -v ./...

clean:
	@echo "Cleaning..."
	@rm -rf \$(BUILD_DIR)

docker-build:
	@echo "Building Docker image..."
	@docker build -t \$(APP_NAME):latest -f build/docker/Dockerfile .

docker-run:
	@echo "Running Docker container..."
	@docker-compose -f deployments/docker-compose/docker-compose.yml up
EOF

echo ""
echo "✅ Go project structure created successfully!"
echo ""
echo "Project: $PROJECT_NAME"
echo "Location: $(pwd)"
echo ""
echo "Next steps:"
echo "  cd $PROJECT_NAME"
echo "  go mod tidy"
echo "  make build"
echo "  make run"
echo ""