package api

import (
	"fmt"
	"github.com/edinpeter/graphql"
	"../game"
)

func NewGetPlayersQueryField(players *game.Players) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(graphql.String),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return players.GetPlayerSummaries(), nil
		},
	}
}

func NewAddPlayerQueryField(players *game.Players) *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"newPlayerName": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"newPlayerGoldCount": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			newPlayerName := p.Args["newPlayerName"].(string)
			newPlayerGoldCount := p.Args["newPlayerGoldCount"].(int)
			playerCount, ok := players.AddPlayer(newPlayerName, newPlayerGoldCount)
			if !ok {
				return "Player name already exists.", nil
			}
			return fmt.Sprintf("New player count: %d", playerCount), nil
		},
	}
}