# go-banking-api

> An api that simulates banking operations and provides functionalities such as:

- Account opening for a new customer
- Making a deposit/withdrawal transaction
- Role based access control

## Third-party Packages Used

- [mysql driver](https://github.com/go-sql-driver/mysql): MySQL driver for Go's (golang) database/sql package
- [sqlx](https://github.com/jmoiron/sqlx): General purpose extensions to golang's database/sql.
- [gorilla mux](https://github.com/gorilla/mux): A powerful HTTP router and URL matcher for building Go web servers.
- [zap](https://github.com/uber-go/zap): Blazing fast, structured, leveled logging in Go.

## Run App

```
SERVER_ADDRESS=localhost SERVER_PORT=8000 DB_USER=root DB_PASSWORD=password DB_HOST=localhost DB_PORT=3306 DB_NAME=go_banking_api go run main.go
```