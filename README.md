# postgres-go-echo-restful

A small skeleton project for a RESTful API using Go with Echo that uses a PostgreSQL database under the hood

## Requirements

- [Git](https://git-scm.com/)
- [Go](https://go.dev/) >= 1.22

## Initial setup

```shell
git clone git@github.com/neowinx/postgres-go-echo-htmx-bulma.git
cd postgres-go-echo-htmx-bulma
go build
```

## How to use

Once you've build the project just run the generated `server` executable. But first you need to define
the required DATABASE_URI env variable

```shell
export DATABASE_URI="postgresql://postgres:postgres@localhost:5432/super"
./server
```

or if you just want to start develop the project. Run the web/server.go file directly

```shell
export DATABASE_URI="postgresql://postgres:postgres@localhost:5432/super"
go run ./web/server.go
```

## SQLC generation

This project uses [sqlc](https://sqlc.dev/) for the database abstraction. In order to generate the queries jsut edit the `querys.sql` file and then execute: 

```shell
sqlc generate
```

to generate the interfaces and magic stuff in `pkg/db`

## Swagger Generation

The swagger documentation is generated using [swaggo](https://github.com/swaggo/swag) and [swagger-echo](https://github.com/swaggo/echo-swagger) respectively. This works by
issuing the next command in order to generate the documentation based on the comments in the go files:

```bash
swag init -g server.go -d ./web,./internal/handler
```

NOTE: This needs to be executed everytime you need to update the swagger documentaion
