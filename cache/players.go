package cache

import "github.com/decendgame/bot/model"

func AddPlayer(player model.Player) {
	ActivePlayers[player.Discord.ID] = &player
}

func GetPlayer(playerID string) (player *model.Player, exists bool) {
	player, exists = ActivePlayers[playerID]
	return
}
