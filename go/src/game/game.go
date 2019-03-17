package game

type GameInstance struct {
	Score int
	GameName string
	TurnCounter int
	PlayerRoster Players 
}

func (game *GameInstance) SetGameName(name string) string {
	game.GameName = name
	return game.GameName
}

func (game *GameInstance) GetGameName() string {
	return game.GameName
}

func (game * GameInstance) IncreaseGameScore(increaseAmount int) int {
	game.Score += increaseAmount
	return game.Score
}

func (game *GameInstance) IncrementTurnCounter() int {
	game.TurnCounter++
	return game.TurnCounter
}

func NewGame(playerCount int, gameName string) GameInstance {
	return GameInstance{0, gameName, 0, NewPlayersRoster() }
}