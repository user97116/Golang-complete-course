package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"hello": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Hello, World!", nil
				},
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: rootQuery,
	},
)

func main() {
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
