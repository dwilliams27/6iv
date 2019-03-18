package main

import (
	"./api"
	"./game"
	"github.com/edinpeter/graphql"
	"github.com/edinpeter/handler"
	"net/http"
)

func main() {
	mainGame := game.NewGame(4, "yeet")
	gamePlayers := mainGame.PlayerRoster

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getNameQuery": api.NewGetNameQueryField(&mainGame),
			"getPlayersQuery": api.NewGetPlayersQueryField(&gamePlayers),
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"putNameQuery": api.NewPutNameQueryField(&mainGame),
			"incrementScoreQuery": api.NewIncreaseScoreQueryField(&mainGame),
			"incrementTurnQuery": api.NewIncrementTurnQueryField(&mainGame),
			"addPlayerQuery": api.NewAddPlayerQueryField(&gamePlayers),
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
