package model

import (
	"github.com/bwmarrin/discordgo"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type Player struct {
	Wallet               hdwallet.Wallet
	Discord              discordgo.User
	Position             int
	Playing              bool
	ActualHouse          House
	IsHerTurn            bool
	PurchaseOfferPending bool
}

func (p *Player) MovePlayer(numberOfHouses int, totalHousesVilla int) (newPos int, completeALap bool) {
	newPos = p.Position + numberOfHouses
	if newPos > totalHousesVilla {
		newPos = totalHousesVilla - newPos
		completeALap = true
	}
	p.Position = newPos
	return
}
