package bot

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/decendgame/bot/cache"
	"github.com/decendgame/bot/config"
	gameCrypto "github.com/decendgame/bot/crypto"
	"github.com/decendgame/bot/model"
	"github.com/decendgame/bot/services/tatum"
)

func GetAnswer(msgRec string, player *model.Player, txID string, session *discordgo.Session, originalMsg *discordgo.MessageCreate) (resp []string) {
	msgRec = strings.ToLower(msgRec)
	msgRec = strings.TrimSpace(msgRec)
	// fmt.Println(msgRec)
	if strings.Contains(msgRec, "hello") || msgRec == "hey" || msgRec == "yo" || msgRec == "hi" {
		resp = initialMsg(txID, *player)
	} else if msgRec == "yes" && player.IsHerTurn && player.PurchaseOfferPending {
		txID, err := gameCrypto.PayHousePurchase(player, config.BotPlayer, cache.Villa[player.Position].GetPrice())
		resp = append(resp, "Payment transaction has been sent. Its ID is "+txID)
		if err != nil {
			resp = append(resp, err.Error())
		} else {
			txID, err = gameCrypto.HouseOwnershipTransfer(player, config.BotPlayer, player.Position)
			if err != nil {
				resp = append(resp, err.Error())
			} else {
				resp = append(resp, "Transfer transaction has been sent. Its ID is "+txID)
			}
		}
	} else if strings.Contains(msgRec, "who are you") {
		respTmp := welcome()
		resp = append(resp, respTmp)
	} else if strings.Contains(msgRec, "story of the game") {
		respTmp := story()
		resp = append(resp, respTmp)
	} else if strings.Contains(msgRec, "how to play") || msgRec == "instructions" {
		respTmp := howToPlay()
		resp = append(resp, respTmp)
	} else if strings.Contains(msgRec, "welcome") {
		respTmp := greetings(player.Discord.Username)
		resp = append(resp, respTmp)
	} else if strings.Contains(msgRec, "start") {
		var err error
		resp, err = startPlayer(player, session, originalMsg)
		if err != nil {
			resp = append(resp, err.Error())
		}
	} else if strings.Contains(msgRec, "list command") {
		respTmp := listCommands()
		resp = append(resp, respTmp)
	} else if strings.Contains(msgRec, "villa map") {
		resp = villaMap()
	}
	// fmt.Println(resp)
	return
}

func startPlayer(player *model.Player, session *discordgo.Session, originalMsg *discordgo.MessageCreate) (msgs []string, err error) {
	err = movePlayer(player)
	if err != nil {
		return
	}
	player.IsHerTurn = true
	msgs = prepareArrival(player, originalMsg)
	return
}

func movePlayer(player *model.Player) error {
	tmp, err := tatum.GetLastestBlock()
	if err != nil {
		return err
	}
	houses, err := strconv.Atoi(tmp)
	if err != nil {
		return err
	}
	player.MovePlayer(houses, cache.NumberOfHouses)
	player.ActualHouse = cache.Villa[player.Position]
	return nil
}

func prepareArrival(p *model.Player, originalMsg *discordgo.MessageCreate) (msg []string) {
	var m string
	m = "You are at " + cache.NFTs[p.ActualHouse.IdOnBoard].Name
	msg = append(msg, m)
	m = "See its NFT in here: https://hoarse-well-made-theemim.explorer.hackathon.skalenodes.com/token/0x2e7c3adc52b963eb108e0273fb7d9b07f8ac6fc1/instance/" + strconv.Itoa(p.ActualHouse.IdOnBoard)
	msg = append(msg, m)
	if cache.Villa[p.ActualHouse.IdOnBoard].Owner.ID != originalMsg.Author.ID {
		if cache.Villa[p.ActualHouse.IdOnBoard].Owner.ID == BotId {
			m = "It is avaliable for purchase. Its price is " + strconv.Itoa(p.ActualHouse.GetPrice())
			msg = append(msg, m)
			m = "Do you want to buy it ? Yes or No?"
			msg = append(msg, m)
			p.PurchaseOfferPending = true
			return
		}
		m = "You are going to pay a rental to " + cache.Villa[p.ActualHouse.IdOnBoard].Owner.Username
		msg = append(msg, m)
		m = "Rental price is: " + strconv.Itoa(cache.Villa[p.ActualHouse.IdOnBoard].RentPrice())
		msg = append(msg, m)
		m = "I am going to perform the transaction on chain..."
		msg = append(msg, m)
	} else {
		m = "Welcome back home, Magister."
		msg = append(msg, m)
	}
	return
}
