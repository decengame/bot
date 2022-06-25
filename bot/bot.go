package bot

import (
	"fmt" // to print errors

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
	var player model.Player
	var exists bool
	var err error
	var returnMessage string

	player, exists = cache.GetPlayer(m.Author.ID)
	if !exists {
		player.Discord = *m.Author
		player.Wallet, err = tatum.CreateWallet()
		if err != nil {
			fmt.Printf("Error creating wallet: %+v\n", err)
			returnMessage = "Oops... en error: " + err.Error()
			_, _ = s.ChannelMessageSend(m.ChannelID, returnMessage)
			return
		}
	}
	// If we message ping to our bot in our discord it will return us pong .
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
