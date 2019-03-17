package api

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"../game"
)

type Schemas struct {
	GqlApiObjectConfigs []graphql.ObjectConfig
}

func NewGetNameQueryField(game *game.GameInstance) *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return game.GetGameName(), nil
		},
	}
}

func NewPutNameQueryField(game *game.GameInstance) *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			newGameName := p.Args["name"].(string)
			game.SetGameName(newGameName)
			return fmt.Sprintf("Changed game name to: %s", game.GameName), nil
		},
	}
}

func NewIncreaseScoreQueryField(game *game.GameInstance) *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"increaseAmount": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			increaseAmount := p.Args["increaseAmount"].(int)
			return fmt.Sprintf("Increased game score to: %d", game.IncreaseGameScore(increaseAmount)), nil
		},
	}
}

func NewIncrementTurnQueryField(game *game.GameInstance) *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return fmt.Sprintf("Incremented game turn counter to: %d", game.IncrementTurnCounter()), nil
		},
	}	
}