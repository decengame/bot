package main

import (
	"fmt" // we will create this later
	"net/http"

	"github.com/decendgame/bot/bot"
	"github.com/decendgame/bot/cache"
	"github.com/decendgame/bot/config"
	"gopkg.in/macaron.v1"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()

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
	m.Run()
	//<-make(chan struct{})
	return
}
