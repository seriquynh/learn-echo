# Learn Echo

This is a sample containing what I learn about Echo, a web framework for Go. In order to build a RESTful API application with MongoDB.

## Requirements

- Go ^1.20
- MongoDB ^8.0

## Installation

Clone the repository.

```bash
cd /path/to/workspace

git clone https://github.com/seriquynh/learn-echo.git

cd learn-echo
```

Run MongoDB service using docker.

```bash
docker compose -f ./docker/docker-compose.yml -f ./docker/docker-compose.local.yml up -d
```

Run the program.

```bash
go run .
```

## Usage

The base URL is `http://127.0.0.1:1323`

| Method | Endpoint         | Description               |
|--------|------------------|---------------------------|
| POST   | /api/users        | Create a new user         |
| GET    | /api/users/:user  | Get user by ID            |
