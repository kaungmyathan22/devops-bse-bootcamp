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

## CI/CD

This project includes GitHub Actions workflow for automated Docker image building and pushing to Docker Hub.

### GitHub Actions Setup

1. Go to your GitHub repository settings
2. Navigate to **Secrets and variables** → **Actions**
3. Add the following secrets:
   - `DOCKER_USERNAME`: Your Docker Hub username (kaungmyathan)
   - `DOCKER_PASSWORD`: Your Docker Hub access token or password

### Workflow Behavior

- **On push to main/master**: Builds and pushes Docker image to `kaungmyathan/devops-bse-bootcamp`
- **On pull requests**: Builds the image but doesn't push (for testing)
- **On tags (v*)**: Creates versioned tags (e.g., `v1.0.0`, `v1.0`, `v1`)
- **Default tag**: `latest` is applied to the main branch

### Pull Docker Image

After the workflow runs, you can pull the image:

```bash
docker pull kaungmyathan/devops-bse-bootcamp:latest
docker run -p 8080:8080 kaungmyathan/devops-bse-bootcamp:latest
```

## Project Structure

```
.
├── .github/
│   └── workflows/
│       └── docker.yml    # GitHub CI/CD workflow
├── main.go               # Main application file
├── go.mod                # Go module file
├── Dockerfile            # Docker configuration
├── .dockerignore         # Docker ignore file
└── README.md             # This file
```

## Building for Production

```bash
# Build binary
go build -o main .

# Run binary
./main
```
