# user-crud

## Project structure
### `cmd/app/main.go`
Configuration and logger initialization.
### `config`
Configuration. First, `config.yml` is read, then environment variables overwrite the yaml config if they match.
### `internal/app`
This is where all the main objects are created.
### `internal/model`
Entities of business logic (models) can be used in any layer.
There can also be methods, for example, for validation.
### `internal/service`
Business logic.
- Methods are grouped by area of application (on a common basis)
- Each group has its own structure
- One file - one structure

### `pkg/httpserver`
- http server with configuration
- ### `pkg/logger`
- standard logger

## Dependency Injection
In order to remove the dependence of business logic on external packages, dependency injection is used.

## examples
```shell
curl -XPOST 'localhost:8080/v1/users' -d @post.json
curl -XGET 'localhost:8080/v1/users'
curl -XPUT 'localhost:8080/v1/users' -d @put.json
```