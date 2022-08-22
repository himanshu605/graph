package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/himanshu605/graphql-api/graph"
	"github.com/himanshu605/graphql-api/graph/generated"
	"github.com/himanshu605/graphql-api/graph/model"
	"github.com/jinzhu/gorm"
)

const defaultPort = "8000"

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:password@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.LogMode(true)

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("CREATE DATABASE test_db")
	db.Exec("USE test_db")

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&model.User{})
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	initDB()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
