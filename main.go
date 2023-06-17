package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: NewResolver()}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
