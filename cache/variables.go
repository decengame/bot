package cache

import "github.com/decendgame/bot/model"

// TATUM-API-KEY the TATUM API KEY
var TATUM_API_KEY string

// ActivePlayers the active players running this game session
var ActivePlayers map[string]model.Player
