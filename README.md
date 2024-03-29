# postgres-go-echo-htmx-bulma

A small skeleton project using Go with Echo and BULMA that uses a PostgreSQL database under the hood

## Swagger Generation

The swagger documentation is generated using swaggo and swagger-echo respectively. This works by
issuing the next command in order to generate the documentation based on the comments in the go
files:

```bash
swag init -g server.go -d ./web,./internal/handler
```

NOTE: This needs to be execute everytime you need to update the swagger documentaion
