package game

import (
	"fmt"
)

type Player struct {
	Name string
	Gold int
}

type Players struct{
	roster map[string]Player
}

func (player *Player) PlayerSummary() map[string]string {
	summary := make(map[string]string)
	summary["PlayerName"] = player.Name
	summary["PlayerGold"] = fmt.Sprintf("%d", player.Gold)
	return summary
}

func (players *Players) AdjustPlayerGold(name string, amount int) int {
	p := players.roster[name]
	p.Gold += amount
	return players.roster[name].Gold
}

func (players *Players) AddPlayer(name string, gold int) (int, bool) {
	if _, ok := players.roster[name]; ok {
		return -1, false
	} else {
    	newPlayer := Player{name, gold}
    	players.roster[name] = newPlayer
    	return len(players.roster), true
    }
}

func (players *Players) GetPlayerSummaries() []map[string]string {
	var playerNames []map[string]string

	for _, player := range players.roster{
		playerNames = append(playerNames, player.PlayerSummary())
	}
	
	return playerNames
}

func NewPlayersRoster() Players {
	return Players{make(map[string]Player)}
}