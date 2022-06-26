package cache

import (
	"fmt"

	"github.com/decendgame/bot/model"
)

func AddPlayer(player model.Player) {
	ActivePlayers[player.Discord.ID] = &player
}

func GetPlayer(playerID string) (player *model.Player, exists bool) {
	player, exists = ActivePlayers[playerID]
	if exists {
		fmt.Print("Player ", player.Discord.Username, " is back\n\n")
	} else {
		fmt.Print("Player ", playerID, " is new\n\n")
	}
	return
}
