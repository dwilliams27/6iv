package game

type GameInstance struct {
	Score int
	GameName string
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