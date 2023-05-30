package main

import (
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/sushanpth/learn-go/go-graphql/schema"
)

func main() {
	h := handler.New(&handler.Config{
		Schema:   &schema.BeastSchema,
		Pretty:   true,
		GraphiQL: false,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
