package cache

import "github.com/decendgame/bot/model"

func init() {
	TATUM_API_KEY = "cedbe61a-44ec-4c8c-b1bb-a481a2534f42"
	ActivePlayers = make(map[string]model.Player)
}
