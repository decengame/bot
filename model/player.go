package model

import (
	"github.com/bwmarrin/discordgo"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type Player struct {
	Wallet  hdwallet.Wallet
	Discord discordgo.User
}
