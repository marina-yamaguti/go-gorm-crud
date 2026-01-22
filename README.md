# Go GORM CRUD API

A RESTful API for managing products built with Go, Gin, GORM, and PostgreSQL. Features automatic API documentation with Swagger/OpenAPI.

## Features

- Complete CRUD operations for products
- PostgreSQL database with GORM
- RESTful API design
- Swagger/OpenAPI documentation
- Docker and Docker Compose support

## Tech Stack

- **Go** - Programming language
- **Gin** - HTTP web framework
- **GORM** - ORM library
- **PostgreSQL** - Database
- **Swagger** - API documentation
- **Docker** - Containerization

## Prerequisites

- Go 1.25.5 or higher
- Docker and Docker Compose (for containerized setup)

## Installation

### Using Docker (Recommended)

1. Clone the repository:

```bash
git clone https://github.com/marina-yamaguti/go-gorm-crud.git
cd go-gorm-crud
```

2. Start the application with Docker Compose:

```bash
docker-compose up -d
```

The API documentationwill be available at:

```bash
http://localhost:8080/swagger/index.html
```
