package model

import (
	"crypto/ecdsa"

	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type Player struct {
	Wallet               hdwallet.Wallet
	Key                  *ecdsa.PrivateKey
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

func (p *Player) EthereumAddress() (address common.Address) {
	publicKey := p.Key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return
	}
	address = crypto.PubkeyToAddress(*publicKeyECDSA)
	return
}
