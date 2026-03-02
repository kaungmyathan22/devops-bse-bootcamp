# Go Web Application

A simple, production-ready Go web application with Docker support.

## Features

- HTTP web server with multiple endpoints
- Health check endpoint (`/health`)
- API endpoint (`/api/hello`)
- Docker containerization
- Multi-stage Docker build for optimized image size

## Prerequisites

- Go 1.21 or higher
- Docker (optional, for containerization)

## Local Development

### Run locally

```bash
# Install dependencies
go mod download

# Run the application
go run main.go
```

The server will start on `http://localhost:8080` by default.

### Available Endpoints

- `GET /` - Home page with information about the app
- `GET /health` - Health check endpoint (returns JSON)
- `GET /api/hello` - API hello endpoint (returns JSON)

### Environment Variables

- `PORT` - Server port (default: 8080)

## Docker

### Build Docker Image

```bash
docker build -t devops-bse-bootcamp .
```

### Run Docker Container

```bash
docker run -p 8080:8080 devops-bse-bootcamp
```

Or with custom port:

```bash
docker run -p 3000:3000 -e PORT=3000 devops-bse-bootcamp
```

### Docker Compose (Optional)

You can create a `docker-compose.yml` file:

```yaml
version: '3.8'
services:
  web:
    build: .
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
```

Then run:

```bash
docker-compose up
```

## Project Structure

```
.
├── main.go          # Main application file
├── go.mod           # Go module file
├── Dockerfile       # Docker configuration
├── .dockerignore    # Docker ignore file
└── README.md        # This file
```

## Building for Production

```bash
# Build binary
go build -o main .

# Run binary
./main
```
