package main

import (
	"./api"
	"./game"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
)

type inputQuery struct {
	age  int
	name string
}

func main() {
	mainGame := game.GameInstance{0, "yeet"}

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getNameQuery": api.NewGetNameQueryObject(&mainGame),
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"putNameQuery": api.NewPutNameQueryObject(&mainGame),
			"incrementScoreQuery": api.NewIncrementScoreQuery(&mainGame),
		},
	})

	queryMutationSchema := graphql.SchemaConfig{Query: queryType, Mutation: mutationType}

	schema, _ := graphql.NewSchema(queryMutationSchema)

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
