# go-fiber-crud-clean-arch

Golang REST API service using Fiber framework and GORM with ProstgresSQL database by applying clean architecture and dependency injection by Wire

## Template Structure

### API

* [Fiber](https://github.com/gofiber/fiber) is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for fast development with zero memory allocation and performance in mind.
* [JWT](github.com/golang-jwt/jwt) A go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens.
* [GORM](https://gorm.io/index.html) with [PostgresSQL](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)The fantastic ORM library for Golang aims to be developer friendly.
* [Wire](https://github.com/google/wire) is a code generation tool that automates connecting components using dependency injection.
* [validator](github.com/go-playground/validator) is a package validator implements value validations for structs and individual fields based on tags.
* [Viper](https://github.com/spf13/viper) is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats.

## Using `go-fiber-crud-clean-arch` project

### Prerequisite

1. Set up your local PostgresSQL database
2. Set up env value which required env key is

* DB_HOST=
* DB_NAME=
* DB_USER=
* DB_PORT=
* DB_PASSWORD=

3. Additional env key is

* RECOVER - to enable the reocery mode for Fiber framework
* TRACING - to enable the log tracing mode for Fiber framework
* PREFORK - to enable use of the SO_REUSEPORT socket option. This socket option allows multiple sockets to listen on the same IP address and port combination. The kernel then load balances incoming connections across the sockets.

### Run application

To use `go-fiber-crud-clean-arch` project, follow these steps:

```bash
# Install dependencies
make deps

# Generate wire_gen.go for dependency injection
# Please make sure you are export the env for GOPATH
make wire

# Run the project in Development Mode
make run
```

Additional commands:

```bash
➔ make help
build                          Compile the code, build Executable File
run                            Start application
test                           Run tests
test-coverage                  Run tests and generate coverage file
deps                           Install dependencies
deps-cleancache                Clear cache in Go module
wire                           Generate wire_gen.go
mockery                        Generate mock file
help                           Display this help screen
```

## Available Endpoint

In the project directory, you can call:

### `GET /healthcheck`

* For getting status page

### `GET /login`

* For generating a JWT

### `GET /api/users`

* For getting list of users

### `GET /api/users/:id`

* For getting user by ID

### `POST /api/users`

* For creating new user

### `DELETE /api/users/:id`

* For removing existing user

### `PUT /api/users/:id`

* For updating existing user

### `GET /api/users/name/:text`

* For retrieving a list of user information that their
name match or partially match with the specified text.

## Folder Structure

This project design by using clean architecture and hexagonal architecture so folder of project will organize base on
clean architecture below

Ref: <https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html>

Here below is how folder map to layer and component in clean architecture

* domain -> Entity
* usecase -> Usecase
* repository -> Repository
* api -> Handler
* driver -> remote call
