package bot

import (
	"fmt" // to print errors

	"github.com/ethereum/go-ethereum/crypto"

	botCrypto "github.com/decendgame/bot/crypto"

	// importing our config package which we have created above
	"github.com/decendgame/bot/cache"
	"github.com/decendgame/bot/config"
	"github.com/decendgame/bot/model"
	"github.com/decendgame/bot/services/tatum"

	"github.com/bwmarrin/discordgo" // discordgo package from the repo of bwmarrin .
)

var (
	BotId string
	goBot *discordgo.Session
)

func Start() {
	// creating new bot session
	goBot, err := discordgo.New("Bot " + config.Token)
	// Handling error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Making our bot a user using User function .
	u, err := goBot.User("@me")
	// Handlinf error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Storing our id from u to BotId .
	BotId = u.ID
	fmt.Println("BotID on Discord is", BotId)

	// Adding handler function to handle our messages using AddHandler from discordgo package. We will declare messageHandler function later.
	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	// Error handling
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// If every thing works fine we will be printing this.
	fmt.Println("Bot is running !")
}

// Definition of messageHandler function it takes two arguments first one is discordgo.Session which is s , second one is discordgo.MessageCreate which is m.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// fmt.Printf("Session: %+v\n\n", s)
	// fmt.Printf("Message: %+v\n\n\n", m)
	// Bot musn't reply to it's own messages , to confirm it we perform this check.
	if m.Author.ID == BotId {
		return
	}
	var player *model.Player
	var exists bool
	var err error
	var txID string

	player, exists = cache.GetPlayer(m.Author.ID)
	if !exists {
		player = new(model.Player)
		player.Discord = *m.Author
		player.Wallet, err = tatum.CreateWallet()
		if err != nil {
			fmt.Printf("Error creating wallet: %+v\n", err)
			returnMessage := "Oops... en error: " + err.Error()
			_, _ = s.ChannelMessageSend(m.ChannelID, returnMessage)
			return
		}
		key, err := crypto.GenerateKey()
		if err != nil {
			fmt.Printf("Error creating private key: %+v\n", err)
			returnMessage := "Oops... en error: " + err.Error()
			_, _ = s.ChannelMessageSend(m.ChannelID, returnMessage)
			return
		}
		player.Key = key
		txID, err = botCrypto.InitalTopup(player.Key, 2000000)
		if err != nil {
			fmt.Printf("Error creating topup: %+v\n", err)
			returnMessage := "Oops... en error: " + err.Error()
			_, _ = s.ChannelMessageSend(m.ChannelID, returnMessage)
			return
		}
		cache.AddPlayer(*player)
	}
	msgs := getAnswer(m.Content, player, s, m)
	for i := 0; i < len(msgs); i++ {
		_, err = s.ChannelMessageSend(m.ChannelID, msgs[i])
		if err != nil {
			fmt.Println("Error sending message:", err.Error())
		}
	}
	if len(txID) > 10 {
		_, err = s.ChannelMessageSend(m.ChannelID, "Transaction ID "+txID+" has been sent to load your account.")
		if err != nil {
			fmt.Println("Error sending message:", err.Error())
		}
		_, err = s.ChannelMessageSend(m.ChannelID, "You don't need to thank me. I know I am good and generous Bot.")
		if err != nil {
			fmt.Println("Error sending message:", err.Error())
		}
		_, err = s.ChannelMessageSend(m.ChannelID, "You can check transaction processing here: ")
		if err != nil {
			fmt.Println("Error sending message:", err.Error())
		}
		_, err = s.ChannelMessageSend(m.ChannelID, "https://hoarse-well-made-theemim.explorer.hackathon.skalenodes.com/tx/"+txID+"/internal-transactions")
		if err != nil {
			fmt.Println("Error sending message:", err.Error())
		}
		tmp := player.EthereumAddress()
		_, err = s.ChannelMessageSend(m.ChannelID, "BTW, your game account wallet is: "+tmp.String())
		if err != nil {
			fmt.Println("Error sending message:", err.Error())
		}
	}
}
