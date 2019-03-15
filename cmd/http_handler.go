package main

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
)

type inputQuery struct {
	age  int
	name string
}

func main() {
	myName := "peter"
	getFields := graphql.Fields{
		"name": &graphql.Field{Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return myName, nil
			},
		},
	}

	putFields := graphql.Fields{
		"changeName": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				myName = p.Args["name"].(string)
				return "changed the name", nil
			},
		},
	}
	getNameQuery := graphql.ObjectConfig{Name: "getNameQuery", Fields: getFields}
	putNameQuery := graphql.ObjectConfig{Name: "putNameQuery", Fields: putFields}

	bothSchemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(getNameQuery), Mutation: graphql.NewObject(putNameQuery)}
	schema, _ := graphql.NewSchema(bothSchemaConfig)

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
