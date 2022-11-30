# Cake API Service

## How Ro Run?

### Using docker compose
1. **Up docker compose**

This will create `mysql` container for development with default database, and build docker image for `cake-api-service`.
```shell
docker-compose up
```
> Before any API is called run init migration first to create
required tables

2. **Init database scheme**

```shell
go run migrations/init.go migrations/mysql/cake_service/001-init_scheme.sql
```

## Exposed API

- `GET /api/v1/cakes` get list of cakes
- `GET /api/v1/cakes/:cakeID` get cake item by id
- `POST /api/v1/cakes` create new cake data
- `PATCH /api/v1/cakes/:cakeID` update cake, one or all fields
- `DELETE /api/v1/cakes/:cakeID` delete cake item by id

## Package Structure
- `config` for accessing external configuration
- `core` is for main domain entity, business logic or use case and interfaces for repository,
    this package will be the center of the structure where other package can depends on and implement the interface
- `handler` is the starting point for http endpoint, it parses the request, pass to core module to process and then return the response
- `repository` will implement repository interface declared in `core`
- `utils` miscellaneous helper functions
- `migrations` contains migration sql file