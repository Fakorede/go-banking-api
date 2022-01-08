package app

import (
	"fmt"
	"go-banking-api/domain"
	"go-banking-api/handlers"
	"go-banking-api/services"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

func checkVars() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASSWORD") == "" ||
		os.Getenv("DB_HOST") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_NAME") == "" {
		log.Fatal("environment variables not defined!!!")
	}
}

func getDBClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	conn, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	return conn
}

func Start() {
	// env vars
	checkVars()

	// db connection
	dbClient := getDBClient()

	// app dependencies
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)

	customerHandlers := handlers.NewCustomerHandlers(services.NewDefaultCustomerService(customerRepositoryDB))
	accountHandlers := handlers.NewAccountHandlers(services.NewDefaultAccountService(accountRepositoryDB))

	// app router
	router := routes(customerHandlers, accountHandlers)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to banking api!")
}
