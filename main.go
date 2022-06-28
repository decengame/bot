package main

import (
	"fmt" // we will create this later
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/decendgame/bot/bot"
	"github.com/decendgame/bot/cache"
	"github.com/decendgame/bot/config"
	"github.com/decendgame/bot/model"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/macaron.v1"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// bot.Start()

	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Get("/tokens/:token.json", func(ctx *macaron.Context) {
		nft, found := cache.NFTs[ctx.ParamsInt("token")]
		if !found {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.JSON(http.StatusOK, nft)
	})
	m.Get("/bot/:text", func(ctx *macaron.Context) {
		BotPlayer := new(model.Player)
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			return
		}
		GodMasterBot := new(discordgo.User)
		GodMasterBot.Email = "jeffprestes@gmail.com"
		GodMasterBot.Username = "DecendGameBot"
		GodMasterBot.Discriminator = "1424"
		GodMasterBot.ID = "990059486054064128"
		BotPlayer.Discord = *GodMasterBot
		BotPlayer.Key = privateKey
		msgs := bot.GetAnswer(ctx.ParamsEscape("text"), BotPlayer, "", nil, nil)
		var sb strings.Builder
		for i := 0; i < len(msgs); i++ {
			sb.WriteString(msgs[i])
		}
		ctx.Status(http.StatusOK)
		ctx.Write([]byte(sb.String()))
		// ctx.HTMLString(http.StatusOK, sb.String())
	})
	m.Run()
	//<-make(chan struct{})
	return
}
