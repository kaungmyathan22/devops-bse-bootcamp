# Go Web Application

A production-ready Go web application with Docker containerization and CI/CD automation.

## Architecture

This application consists of:
- **Go HTTP Server**: Lightweight web server with health checks, API endpoints, and metrics
- **Docker**: Multi-stage build producing optimized Alpine-based container images
- **CI/CD**: GitHub Actions workflow that automatically builds and pushes images to Docker Hub on every push to main

## Features

- HTTP web server with multiple endpoints
- Production-ready health check endpoint with uptime and version tracking
- Metrics endpoint for observability
- Docker containerization with non-root user
- Automated CI/CD pipeline
- Comprehensive test coverage

## Running Locally

### Prerequisites

- Go 1.21 or higher
- Docker (optional, for containerization)

### Setup and Run

```bash
# Install dependencies
go mod download

# Run the application
go run main.go
```

The server will start on `http://localhost:8080` by default.

### Available Endpoints

- `GET /` - Home page with information about the app
- `GET /health` - Health check endpoint (returns status, uptime, version)
- `GET /api/hello` - API hello endpoint
- `GET /metrics` - Metrics endpoint with request counters

### Environment Variables

- `PORT` - Server port (default: 8080)

### Running Tests

```bash
go test ./...
```

### Building with Version

To build with a specific version:

```bash
go build -ldflags "-X main.Version=v1.0.0" -o main .
```

## Running with Docker

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

## CI/CD

This project uses GitHub Actions for automated builds and deployments.

### Setup

1. Go to your GitHub repository settings
2. Navigate to **Secrets and variables** → **Actions**
3. Add the following secrets:
   - `DOCKER_USERNAME`: Your Docker Hub username (kaungmyathan)
   - `DOCKER_PASSWORD`: Your Docker Hub access token or password

### Workflow Behavior

- **Every push to main/master**: Automatically runs tests, builds Docker image, and pushes to Docker Hub as `kaungmyathan/devops-bse-bootcamp:latest`
- **On pull requests**: Runs tests and builds the image but doesn't push (for validation)
- **Tagging vX.Y.Z**: Creates versioned image tags (e.g., `v1.0.0`, `v1.0`, `v1`) in addition to `latest`

### Pull Docker Image

After the workflow runs, you can pull the image:

```bash
docker pull kaungmyathan/devops-bse-bootcamp:latest
docker run -p 8080:8080 kaungmyathan/devops-bse-bootcamp:latest
```

For versioned releases:

```bash
docker pull kaungmyathan/devops-bse-bootcamp:v1.0.0
docker run -p 8080:8080 kaungmyathan/devops-bse-bootcamp:v1.0.0
```

## Project Structure

```
.
├── .github/
│   └── workflows/
│       └── docker.yml    # GitHub CI/CD workflow
├── main.go               # Main application file
├── main_test.go          # Test file
├── go.mod                # Go module file
├── Dockerfile            # Docker configuration
├── .dockerignore         # Docker ignore file
└── README.md             # This file
```

## Building for Production

```bash
# Build binary with version
go build -ldflags "-X main.Version=v1.0.0" -o main .

# Run binary
./main
```
