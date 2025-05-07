package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/theycallmereza/tidymind-backend/graph"
	"github.com/theycallmereza/tidymind-backend/graph/generated"
	"log"
	"net/http"
)

func main() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("Server is running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
