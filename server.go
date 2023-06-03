package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/axeldeveloper/go-gqlgen-todos/graph"
	"github.com/axeldeveloper/go-gqlgen-todos/graph/generated"
	"github.com/labstack/gommon/color"
)

const defaultPort = "8080"

func lineSeparator() {
	fmt.Println("=========================================================")
}

func startMessage() {
	lineSeparator()
	color.Green("Listening on localhost%s\n", defaultPort)
	color.Green("Visit `http://localhost%s/graphql` in your browser\n", defaultPort)
	lineSeparator()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	startMessage()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
