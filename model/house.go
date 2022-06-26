package model

import "github.com/bwmarrin/discordgo"

type House struct {
	IdOnBoard int
	TokenURI  string
	price     int
	Owner     *discordgo.User
}

func NewHouse(idOnBoard int, tokenURI string, owner *discordgo.User) (newHouse House) {
	newHouse.price = 10000
	newHouse.IdOnBoard = idOnBoard
	newHouse.TokenURI = tokenURI
	newHouse.Owner = owner
	return
}

func (h House) RentPrice() (rentPrice int) {
	rentPrice = (h.price * 10) / 100
	return
}

func (h *House) ReadjustPrice() {
	h.price = (h.price * 5) / 100
}

func (h *House) SetPrice(newPrice int) {
	if newPrice < 10000 {
		return
	}
}

func (h House) GetPrice() (price int) {
	price = h.price
	return
}

type HouseNFT struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image"`
	Decimals    int    `json:"decimals"`
}

func NewHouseNFT(name string, description string, imageURL string) (nft HouseNFT) {
	nft.Decimals = 0
	nft.Description = description
	nft.ImageURL = imageURL
	nft.Name = name
	return
}
