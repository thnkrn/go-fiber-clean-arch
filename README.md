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
